/*
 * File  : main.go
 * Author: Lch <l.channng@gmail.com>
 * Date  : 2019/05/16 14:53:48
 */

package main

import (
	"fmt"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}

/*
showA
showB
*/
