package main

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var wg sync.WaitGroup

func main() {
	var engine = Connect()
	defer engine.Close()

	// engine.Sync2(new(User), new(Point)) //=> create table User and Point

	// var insertP = Point{
	// 	User_id:    insertU.Id,
	// 	Points:     10,
	// 	Max_points: 10,
	// }
	// insertUser()
	// insertPoint(&insertP)
	// var readU = User{
	// 	Id:         `4`,
	// 	Name:       `Aguero`,
	// 	Birth:      30,
	// 	Created:    26,
	// 	Updated_at: 231,
	// }

	// var readP = Point{
	// 	User_id:    readU.Id,
	// 	Points:     10,
	// 	Max_points: 10,
	// }
	// readUser(&readU)
	// readPoint(&readP)

	// var updateU = User{
	// 	Id:         `3`,
	// 	Name:       `Dias`,
	// 	Birth:      26,
	// 	Created:    20,
	// 	Updated_at: 23100,
	// }
	// updateUser(&updateU)
	// var p1 = Point{
	// 	User_id:    u1.Id,
	// 	Points:     9,
	// 	Max_points: 10,
	// }
	//updateBirth("2", 100) //=>EX 2 update Birth

	insert100User()
}
