package main

func main() {
	println("beg")

	var stack []string

	stack = append(stack, "world")
	stack = append(stack, "hello")

	for len(stack) > 0 {
		n := len(stack) - 1
		println(stack[n])
		stack = stack[:n]
	}
	println("end")
}
