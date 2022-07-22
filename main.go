package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	
	
)

var Xorm *xorm.Engine


func main() {
	var err = error
	var engine, err := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")

}


type user struct {
	Id   string
	Name string
	birth int64
	created int64
	updated_at int64
}
type point struct {
	user_id string
	points int64
	max_points int64

}

vdff
