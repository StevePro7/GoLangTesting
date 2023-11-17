package main

import "bytes"

func main() {

	sl1 := []byte{'I', 'N', 'T', 'E', 'R', 'V', 'I', 'E', 'W'}
	sl2 := []byte{'B', 'I', 'T'}

	res := bytes.Compare(sl1, sl2)
	println(res)
}
