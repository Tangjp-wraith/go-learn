package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	dumpPrint()
	endTime := time.Now()
	fmt.Println("dumpPrint time cost: " + endTime.Sub(startTime).String())

	startTime = time.Now()
	joinPrint()
	endTime = time.Now()
	fmt.Println("joinPrint time cost: " + endTime.Sub(startTime).String())
}

func dumpPrint() {
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func joinPrint() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
