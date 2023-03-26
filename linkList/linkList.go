/**************************************************************************************************/
/* Copyright (C) mc2lab.com, SSE@USTC, 2022-2023                                                  */
/*                                                                                                */
/*  FILE NAME             :  linkList.go                                                          */
/*  PRINCIPAL AUTHOR      :  WangXin                                                              */
/*  SUBSYSTEM NAME        :  menu                                                                 */
/*  MODULE NAME           :  linkList                                                             */
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

package linkList

import "fmt"

type DataNode struct {
	Cmd     string
	Desc    string
	Handler func() int
	Next    *DataNode
}

func FindCmd(head *DataNode, cmd string) *DataNode {
	for head != nil && head.Cmd != cmd {
		head = head.Next
	}
	return head
}

func ShowAllCmd(head *DataNode) int {
	fmt.Println("Menu List:")
	for head != nil {
		fmt.Printf("%-7s - %s\n", head.Cmd, head.Desc)
		head = head.Next
	}
	return 0
}
