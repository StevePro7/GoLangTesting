package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

// An Item is a stripped down RSS item.
type Item struct {
	Title   string
	Channel string
	GUID    string
}

// A Fetcher fetches Items and returns the time when the next fetch should be
// attempted.  On failure, Fetch retuns a non-nil error.
type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

// A Subscription delivers Items over a channel.  Close cancels the
// subscription, closes the Updates channel, and returns the last fetch error,
// if any.
type Subscription interface {
	Updates() <-chan Item
	Close() error
}

// sub implements the Subscription interface.
type sub struct {
	fetcher Fetcher         // fetches items
	updates chan Item       // sends items to the user
	closing chan chan error // for Close
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc // HLchan
	return <-errc     // HLchan
}

// loop periodically fetches Items, sends them on s.updates, and exits
// when Close is called.  It extends dedupeLoop with logic to run
// Fetch asynchronously.
func (s *sub) loop() {
	const maxPending = 10
	type fetchResult struct {
		fetched []Item
		next    time.Time
		err     error
	}

	var fetchDone chan fetchResult // if non-nil then Fetch is running		// HL
	var pending []Item
	var next time.Time
	var err error
	var seen = make(map[string]bool)

	for {
		var fetchDelay time.Duration
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}

		var startFetch <-chan time.Time
		if fetchDone == nil && len(pending) < maxPending { // HLfetch
			startFetch = time.After(fetchDelay) // enable fetch case
		}

		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // enable send case
		}

		select {
		case <-startFetch: // HLfetch
			fetchDone = make(chan fetchResult, 1) // HLfetch
			go func() {
				fetched, next, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, next, err}
			}()
		case result := <-fetchDone: // HLfetch
			fetchDone = nil // HLfetch
			// Use result.fetched, result.next, result.err
			fetched := result.fetched
			next, err = result.next, result.err
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			for _, item := range fetched {
				if id := item.GUID; !seen[id] { //HLdupe
					pending = append(pending, item)
					seen[id] = true //HLdupe
				}
			}
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
			return
		case updates <- first:
			pending = pending[1:]
		}
	}
}

// Subscript returns a new Subscription that uses fetcher to fetch Items.
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),       // for Updates
		closing: make(chan chan error), // for Close
	}
	go s.loop()
	return s
}

type naiveSub struct {
	fetcher Fetcher
	updates chan Item
	closed  bool
	err     error
}

func (s *naiveSub) Updates() <-chan Item {
	return s.updates
}
func (s *naiveSub) Close() error {
	s.closed = true // HLsync
	return s.err    // HLsync
}
func (s *naiveSub) loop() {
	for {
		if s.closed { // HLsync
			close(s.updates)
			return
		}
		items, next, err := s.fetcher.Fetch()
		if err != nil {
			s.err = err                  // HLsync
			time.Sleep(10 * time.Second) // HLsleep
			continue
		}
		for _, item := range items {
			s.updates <- item // HLsend
		}
		if now := time.Now(); next.After(now) {
			time.Sleep(next.Sub(now)) // HLsleep
		}
	}
}

func NaiveSubscribe(fetcher Fetcher) Subscription {
	s := &naiveSub{
		fetcher: fetcher,
		updates: make(chan Item),
	}
	go s.loop()
	return s
}

type fakeFetcher struct {
	channel string
	items   []Item
}

// FakeDuplicates causes the fake fetcher to return duplicate items.
var FakeDuplicates bool

func (f *fakeFetcher) Fetch() (items []Item, next time.Time, err error) {
	now := time.Now()
	next = now.Add(time.Duration(rand.Intn(5)) * 500 * time.Millisecond)
	item := Item{
		Channel: f.channel,
		Title:   fmt.Sprintf("[%d] Item %d", goid.ID(), len(f.items)),
	}
	item.GUID = item.Channel + "/" + item.Title
	f.items = append(f.items, item)
	if FakeDuplicates {
		items = f.items
	} else {
		items = []Item{item}
	}
	return
}

func fakeFetch(domain string) Fetcher {
	return &fakeFetcher{
		channel: domain,
	}
}

// Fetch returns a Fetcher for Items from domain.
func Fetch(domain string) Fetcher {
	return fakeFetch(domain)
}

type merge struct {
	subs    []Subscription
	updates chan Item
	quit    chan struct{}
	errs    chan error
}

func (m *merge) Updates() <-chan Item {
	return m.updates
}
func (m *merge) Close() (err error) {
	close(m.quit)
	for range m.subs {
		if e := <-m.errs; e != nil { // HL
			err = e
		}
	}
	close(m.updates) // HL
	return
}

// Merge returns a Subscription that merges the item streams from subs.
// Closing the merged subscription closes subs.
func Merge(subs ...Subscription) Subscription {
	m := &merge{
		subs:    subs,
		updates: make(chan Item),
		quit:    make(chan struct{}),
		errs:    make(chan error),
	}

	for _, sub := range subs {
		go func(s Subscription) {
			for {
				var it Item
				select {
				case it = <-s.Updates():
				case <-m.quit: // HL
					m.errs <- s.Close() // HL
					return              // HL
				}
				select {
				case m.updates <- it:
				case <-m.quit: // HL
					m.errs <- s.Close() // HL
					return              // HL
				}
			}
		}(sub)
	}
	return m
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())

	// Subscribe to some feeds, and create a merged update stream.
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com")),
	)

	// Close the subscriptions after some time.
	time.AfterFunc(3*time.Second, func() {
		fmt.Printf("[%d] closed: %v\n", goid.ID(), merged.Close())
	})

	// Print the stream.
	for it := range merged.Updates() {
		fmt.Printf("[%d] %s %s\n", goid.ID(), it.Channel, it.Title)
	}

	fmt.Printf("[%d] end\n", goid.ID())
	panic("show me the stacks")
}
