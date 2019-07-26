/*
 * File  : chan_mutex_test.go
 * Author: Lch <l.channng@gmail.com>
 * Date  : 2019/05/20 19:52:33
 */

package main

import (
    "testing"
)

type ConcreteType struct{
	n int
}

func BenchmarkNoCast(b *testing.B) {
	c := &ConcreteType{}
    for i := 0; i < b.N; i++ {
    	c.n++
    }
}

func BenchmarkDownCast(b *testing.B) {
	var a interface{}
	a = &ConcreteType{}
    for i := 0; i < b.N; i++ {
		c := a.(*ConcreteType)
		c.n++
    }
}
