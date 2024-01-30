package main

import (
	"fmt"
	"gopl/chapter2/popcount"
)

func main() {
	x := uint64(189264826)
	fmt.Printf("count: %d\n", popcount.PopCount(x))
	fmt.Printf("count: %d\n", popcount.LPopCount(x))
	fmt.Printf("count: %d\n", popcount.SPopCount(x))
	fmt.Printf("count: %d\n", popcount.CPopCount(x))
}
