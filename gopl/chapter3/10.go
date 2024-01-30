package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	k := n % 3
	ans := ""
	if k != 0 {
		ans = s[:k] + ","
	}
	for i := k; i < n; i += 3 {
		if i+3 == n {
			ans += s[i : i+3]
		} else {
			ans += s[i:i+3] + ","
		}
	}
	return ans
}

func commaCov(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma("12") == commaCov("12"))
}

func commaAns(s string) string {
	var buf bytes.Buffer
	runeCnt := utf8.RuneCountInString(s) // 先看有几个字符
	sep := runeCnt % 3
	cnt := 0
	for _, r := range s {
		// r is rune
		if cnt%3 == sep && cnt != 0 { // 每个逗号的位置为 sep + 3n
			buf.WriteString(",")
		}
		buf.WriteRune(r)
		cnt++
	}
	return buf.String()
}
