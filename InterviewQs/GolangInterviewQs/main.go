package main

func routine(ch chan struct{}) {
	<-ch
	println("recd_func")
	close(ch)
}

func main() {
	ch := make(chan struct{})
	go routine(ch)
	ch <- struct{}{}
	<-ch
	println("recd_main")
}
