/*
 * File  : chan.go
 * Author: Lch <l.channng@gmail.com>
 * Date  : 2019/07/22 22:29:11
 */

package main

import (
    "log"
)

func main() {
    die := make(chan struct{})

    go func(die chan struct{}) {
        select {
        case <- die:
            log.Println("die 1")
        }
    }(die)

    go func(die chan struct{}) {
        select {
        case <- die:
            log.Println("die 2")
        }
    }(die)

    close(die)

    _, ok := <- die
    if ok {

    }
}