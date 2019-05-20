/*
 * File  : chan_mutex_test.go
 * Author: Lch <l.channng@gmail.com>
 * Date  : 2019/05/20 19:52:33
 */

package main

import (
    "sync"
    "testing"
)

var ch = make(chan bool, 1)
func BenchmarkChan(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ch <- true
        <- ch
    }
}

var mutex = sync.Mutex{}
func BenchmarkMutex(b *testing.B) {
    for i := 0; i < b.N; i++ {
        mutex.Lock()
        mutex.Unlock()
    }
}
