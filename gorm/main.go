/*
 * File  : main.go
 * Author: Lch <l.channng@gmail.com>
 * Date  : 2019/05/14 11:22:35
 */

package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
    Id int64
    Name string
    Age int
}

func main() {
    db, err := gorm.Open("mysql", "root:ijk123@/gorm?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println("connection succedssed")
    }

    defer db.Close()

    db.DropTable(&User{})
    db.AutoMigrate(&User{})

    for i := 0; i < 10; i++ {
        user := User{Name: "lch", Age:20}
        if err := db.Create(&user).Error; err != nil {
            fmt.Println("failec to insert ", err)
        }
    }

    found := &User{Id:3}
    db.Model(found).Update("name", "gorm")
}


