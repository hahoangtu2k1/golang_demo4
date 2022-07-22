package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

var Xorm *xorm.Engine
var engine, err := xorm.NewEngine("mysql", "")

func main() {
	fmt.Println("Hello, world!")

}

type createNew struct {
	Id   int64
	Name string `xorm:"unique"`
}
