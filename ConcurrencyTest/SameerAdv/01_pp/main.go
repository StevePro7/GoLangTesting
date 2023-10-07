package main

import (
	"fmt"

	"github.com/nikandfor/goid"
)

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())
	fmt.Printf("[%d] end\n", goid.ID())
}
