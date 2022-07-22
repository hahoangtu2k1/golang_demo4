package main

import "fmt"

var engine = Connect()

func insertMysql(u *User, p *Point) {
	_, err := engine.Table("user").Insert(&u)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Insert successfully USER")
	}
	_, err = engine.Table("point").Insert(&p)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Insert successfully Point")
	}

}
