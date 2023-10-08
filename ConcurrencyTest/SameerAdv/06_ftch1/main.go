package main

import (
	"fmt"
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

// Subscript returns a new Subscription that uses fetcher to fetch Items.
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{}

	return s
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())

	fmt.Printf("[%d] end\n", goid.ID())
}
