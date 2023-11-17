package main

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	println("beg")
	pos := adder()
	println(pos(1))
	println(pos(3))
	println(pos(6))
	println("end")
}
