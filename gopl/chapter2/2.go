package main

import (
	"fmt"
	"gopl/chapter2/lenconv"
	"os"
	"strconv"
)

func main() {
	var ft lenconv.Foot
	if len(os.Args) == 1 {
		fmt.Println("please input a number (unit ft):")
		var input float64
		fmt.Scanf("%g", &input)
		ft = lenconv.Foot(input)
		fmt.Printf("%s = %s = %s\n", ft, lenconv.FToM(ft).String(), lenconv.FToI(ft).String())

	} else {
		for _, inputStr := range os.Args[1:] {
			f, err := strconv.ParseFloat(inputStr, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err: %v", err)
			}

			ft = lenconv.Foot(f)

			fmt.Printf("%s = %s = %s\n", ft, lenconv.FToM(ft).String(), lenconv.FToI(ft).String())
		}
	}
}
