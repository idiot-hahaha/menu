package main

import "fmt"

func main() {
	var cmd string
	for {
		fmt.Printf("请输入命令:")
		fmt.Scanln(&cmd)
		if cmd == "help" {
			fmt.Println("hello: 打印\"hello\"")
			fmt.Println("exit:  退出")
			fmt.Println("help:  获取命令信息")
		} else if cmd == "hello" {
			fmt.Println("hello")
		} else if cmd == "exit" {
			fmt.Println("退出")
			break
		} else {
			fmt.Println("命令无效,输入help获取命令信息")
		}
	}
}
