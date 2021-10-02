Debugging Go Routine leaks
30/09/2021

https://blog.min.io/debugging-go-routine-leaks



Good practices
01.
blocked on a for-select
for {
	select {
		case <c:
		// process here
	}
}

02.
Using timeout channel
timeout := make(chan bool, 1)
go func() {
	time.Sleep(ie9)	// one second
	timeout <- true
}()

select {
	case <- ch:
		// a read from ch has occurred

	case <- timeout:
		// the read from ch has timed out	
}

OR

select {
	case res := <-c1:
		fmt.Println(res)
	
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
}


context
https://pkg.go.dev/context
