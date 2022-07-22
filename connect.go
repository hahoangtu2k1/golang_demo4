package main

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func Connect() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	// defer engine.Close()

	return engine
}
