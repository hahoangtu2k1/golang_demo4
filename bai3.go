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
func scanTable(id string, get User) *User {
	rows, err := engine.Rows(&get)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
	create := new(User)
	for rows.Next() {
		err = rows.Scan(create)
		if err != nil {
			fmt.Println(err)
		}
	}
	return create
}
func scanTableUser() error {
	countId := make(chan string, 10)
	countName := make(chan string, 10)
	getUser := User{}
	countAll, err := engine.Count(&getUser)
	if err != nil {
		return err
	} else {
		for i := 1; i <= int(countAll); i++ {
			wg.Add(1)
			takeId := strconv.Itoa(i)
			takeAll := scanTable(takeId, User{})
			countId <- takeAll.Id
			countName <- takeAll.Name
			go worker(i, countId, countName)
			wg.Done()
			wg.Wait()
		}
		close(countId)
		close(countName)
	}

	return err

}
func worker(i int, countId <-chan string, countName <-chan string) {
	fmt.Printf("number :%v - ID: %v - Name: %v\n", i, <-countId, <-countName)
}
