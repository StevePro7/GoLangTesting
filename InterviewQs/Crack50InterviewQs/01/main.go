package main

func main() {
	quit := make(chan bool)
	println("beg")
	go func() {
		for {
			select {
			case <-quit:
				println("quitting...")
				return
			}
		}
	}()

	println("about to quit")
	quit <- true
	println("end")
}
