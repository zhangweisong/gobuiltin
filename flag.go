package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func Args() {
	fmt.Println(os.Args)
}

func Flag() {

	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	// flag.Type(flag名, 默认值, 帮助信息)
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println("Args", flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println("NArg", flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println("NFlag", flag.NFlag())
}

// 命令行参数的解析
func main() {
	Flag()
}
