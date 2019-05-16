/*
 * File  : main.go
 * Author: Lch <l.channng@gmail.com>
 * Date  : 2019/05/16 14:53:48
 */

package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("111") }()
	defer func() { fmt.Println("222") }()
	defer func() { fmt.Println("333") }()
	panic("panic")
}
