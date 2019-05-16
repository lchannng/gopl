/*
 * File  : main.go
 * Author: Lch <l.channng@gmail.com>
 * Date  : 2019/05/16 14:53:48
 */

package main

import (
    "fmt"
    "runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}
}

/*
随机panic，select随机公平地选择(???)case执行，后续看下代码下个blog
https://lchannng.github.io/2019/05/16/go-select-implementation-md
*/
