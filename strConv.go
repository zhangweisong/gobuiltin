package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型与其字符串表示的转换

func main() {
	fmt.Println(strconv.Atoi("100"))
	fmt.Println(strconv.Itoa(1125))
	fmt.Println(strconv.ParseInt("-2", 10, 64))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.FormatInt(12, 16))
}
