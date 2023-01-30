package main

import (
	"fmt"
	"log"
	"os"
)

func basic() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}

func flag() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}
func prefix() {
	log.SetPrefix("[pprof]")
	log.Println("这是一条很普通的日志。")
}

func logFile() {
	// 打开文件句柄
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.Println("这是一条很普通的日志。")
}

func instance() {
	Log := log.New(os.Stdout, "[哈哈哈]", os.O_CREATE|os.O_WRONLY|os.O_APPEND)
	Log.Println("dsdsadasd")
}

func main() {
	instance()
}
