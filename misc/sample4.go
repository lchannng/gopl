/*
 * File  : main.go
 * Author: Lch <l.channng@gmail.com>
 * Date  : 2019/05/16 14:53:48
 */

package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    loc1, err := time.LoadLocation("") // utc
    if err != nil {
        fmt.Println(err)
        return
    }

    loc2, err := time.LoadLocation("Local")
    if err != nil {
        fmt.Println(err)
        return
    }

    loc3, err := time.LoadLocation("America/New_York")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(now)
    fmt.Println(now.UTC())
    fmt.Println(now.In(loc1))
    fmt.Println(now.In(loc2))
    fmt.Println(now.In(loc3))
}
