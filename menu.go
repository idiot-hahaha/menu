/**************************************************************************************************/
/* Copyright (C) mc2lab.com, SSE@USTC, 2022-2023                                                  */
/*                                                                                                */
/*  FILE NAME             :  menu.go                                                              */
/*  PRINCIPAL AUTHOR      :  WangXin                                                              */
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
* Created by WangXin, 2023/03/25
*
 */
package main

import (
	"fmt"
	"menu/linkList"
	"os"
)

const (
	MAX_CMD_LENGTH = 128
	DESC_LEN       = 1024
	CMD_NUM        = 10
)

var head = &linkList.DataNode{
	Cmd:     "help",
	Desc:    "this is help command",
	Handler: helpHandler,
	Next: &linkList.DataNode{
		Cmd:     "version",
		Desc:    "menu program v1.0",
		Handler: nil,
		Next: &linkList.DataNode{
			Cmd:     "quit",
			Desc:    "exit the program",
			Handler: quitHandler,
			Next:    nil,
		},
	},
}

func helpHandler() int { return 0 }

func quitHandler() int {
	os.Exit(0)
	return 0
}

func main() {
	head.Handler = func() int {
		linkList.ShowAllCmd(head)
		return 0
	}
	for {
		cmd := make([]byte, MAX_CMD_LENGTH)
		fmt.Printf("Input a cmd number > ")
		fmt.Scanln(&cmd)
		p := linkList.FindCmd(head, string(cmd))
		if p == nil {
			fmt.Printf("This is a wrong cmd!\n ")
			continue
		}
		fmt.Printf("%s - %s\n", p.Cmd, p.Desc)
		if p.Handler != nil {
			p.Handler()
		}
	}
}
