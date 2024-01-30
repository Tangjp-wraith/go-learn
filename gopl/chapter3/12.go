package main

import (
	"fmt"
	"reflect"
)

func stringInMap(str string) map[byte]int {
	count := make(map[byte]int)
	for _, s := range str {
		count[byte(s)]++
	}
	return count
}

func check(s, t string) bool {
	count1 := stringInMap(s)
	count2 := stringInMap(t)
	if reflect.DeepEqual(count1, count2) && s != t {
		return true
	}
	return false
}

func main() {
	fmt.Println(check("adfsd", "123213"))
	fmt.Println(check("adfsd", "adfsd"))
	fmt.Println(check("adfsd", "fsdad"))
}
