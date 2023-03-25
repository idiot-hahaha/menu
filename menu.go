package main

import "fmt"

/**************************************************************************************************/
/* Copyright (C) mc2lab.com, SSE@USTC, 2022-2023                                                  */
/*                                                                                                */
/*  FILE NAME             :  menu.go                                                              */
/*  PRINCIPAL AUTHOR      :  Mengning                                                             */
/*  SUBSYSTEM NAME        :  menu                                                                 */
/*  MODULE NAME           :  menu                                                                 */
/*  LANGUAGE              :  Golang                                                               */
/*  TARGET ENVIRONMENT    :  ANY                                                                  */
/*  DATE OF FIRST RELEASE :  2023/03/25                                                           */
/*  DESCRIPTION           :  This is a menu program                                               */
/**************************************************************************************************/

/*
 * Revision log:
 *
 * Created by Mengning, 2023/03/25
 *
 */

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
