package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func commaAll(s string) string {
	var buf bytes.Buffer
	if s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	dotIndex := strings.LastIndex(s, ".")
	decimals := s[dotIndex:] // 小数部分
	s = s[:dotIndex]
	runeCnt := utf8.RuneCountInString(s) // 先看有几个字符
	sep := runeCnt % 3                   // 距字符数计算余数应为几才是逗号位置
	cnt := 0
	for _, r := range s {
		// r is rune
		if cnt%3 == sep && cnt != 0 { // 每个逗号的位置为 sep + 3n
			buf.WriteString(",")
		}
		buf.WriteRune(r)
		cnt++
	}
	buf.WriteString(decimals)
	return buf.String()
}

func main() {
	fmt.Println(commaAll("-1234.121233"))
}
