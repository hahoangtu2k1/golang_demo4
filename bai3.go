package main

import (
	"fmt"
	"strconv"
)

func insert100User() {
	for i := 5; i <= 105; i++ {
		insertID := strconv.Itoa(i)
		inU := &User{
			Id:         insertID,
			Name:       "UserName_" + strconv.Itoa(i),
			Birth:      int64(i),
			Created:    int64(i),
			Updated_at: int64(i),
		}
		inP := &Point{
			User_id:    insertID,
			Points:     10,
			Max_points: 10,
		}
		_, err := engine.Table("user").Insert(&inU)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("insert", i, " USER successfully")
		}
		_, err = engine.Table("point").Insert(&inP)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("insert", i, "POINT successfully")
		}

	}

}
func scanTable() {

	var get User
	// 	_, err := engine.Table("user").Get(&User{
	// 		Id:   get.Id,
	// 		Name: get.Name,
	// 	})
	// 	if err != nil {
	// 		fmt.Print(err)
	// 	}

	// }
	rows, err := engine.Rows(&get)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
	cre := new(User)
	for rows.Next() {
		err = rows.Scan(cre)
		if err != nil {
			fmt.Println(err)
		}
	}

}
func scanTableUser() error {
	countId := make(chan string, 2)
	countName := make(chan string, 2)
	getUser := User{}
	countAll, err := engine.Count(&getUser)
	if err != nil {
		return err
	} else {
		var g int64
		for g = 1; g <= countAll; g++ {
			wg.Add(1)

		}

	}
}
func worker(i int, countId <-chan string, countName <-chan string) {
	fmt.Printf("count :%v - ID: %v - Name: %v\n", i, <-countId, <-countName)
}
