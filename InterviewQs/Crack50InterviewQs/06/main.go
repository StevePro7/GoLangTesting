package main

func Sum(_, _ chan int) {
}

func main() {
	println("beg")

	inptchannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 0; i < 6; i++ {
			sum += <-inptchannel
		}
		println(sum)
	}()
	Sum(inptchannel, quitchannel)
	println("end")
}
