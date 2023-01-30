package main

import (
	"fmt"
	"io"
	"os"
)

func open() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	len, err := file.WriteString("create\n")
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(len)
	file.Write([]byte("create\n"))
	// 关闭文件
	file.Close()
}

func read() {
	// 打开文件
	file, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}
func create() {
	// 新建文件
	file, err := os.Create("./xxx.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("create\n")
	file.Write([]byte("create\n"))
}

func copy() {
	// 打开源文件
	file1, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新文件
	file2, err2 := os.Create("./abc2.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 缓冲读取
	buf := make([]byte, 1024)
	for {
		// 从源文件读数据
		n, err := file1.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//写出去
		file2.Write(buf[:n])
	}
	file1.Close()
	file2.Close()
}
func main() {
	copy()
}
