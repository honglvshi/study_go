package main

import (
	"fmt"
	"strings"
)

//字符串常用的操作
//其他类型 转换string

func main() {

	data := []byte{100, 101, 98, 117, 103}

	fmt.Println(byteToString(data[:]))

	fmt.Println(stringToByte("debug"))

	fmt.Println(stringCombineByJoin("a", "b", "c"))
	fmt.Println(stringCombineBySymbo("a", "b", "c"))
}

func byteToString(b []byte) string {
	return string(b)
}

func stringToByte(s string) []byte {
	return []byte(s)
}

func stringCombineByJoin(s ...string) string {
	return strings.Join(s, "")
}

func stringCombineBySymbo(s ...string) string {

	content := ""
	for _, v := range s {
		content += v
	}

	return content
}
