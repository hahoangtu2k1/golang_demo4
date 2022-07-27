package main

import (
	"fmt"
)

var engine = Connect()

//insert Mysql

func insertUser(insertU *User) {

	_, err := engine.Table("user").Insert(&insertU)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Insert successfully USER")
	}
}
func insertPoint(insertP *Point) {
	_, err := engine.Table("point").Insert(&insertP)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Insert successfully Point")
	}
}

//read Mysql
// func readUser(readU *User) {
// 	_, err := engine.Table("user").Get(&readU)
// 	if err != nil {
// 		panic(err)
// 	}
// }
// func readPoint(readP *Point) {
// 	_, err := engine.Table("point").Get(&readP)
// 	if err != nil {
// 		panic(err)
// 	}
// }

//update Mysql
func updateUser(updateU *User) {
	// User := new(User)
	_, err := engine.Table("user").Where("id = ?", updateU.Id).Update(updateU)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Update  USER successfully")
	}

}

// func updatePoint(p1 *Point) {
// 	_, err := engine.Table("point").Where("id = ?", p1.User_id).Update(p1)
// 	if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("Update POINT successfully")
// 	}
// }
