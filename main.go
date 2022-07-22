package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var engine = Connect()
	defer engine.Close()
	// engine.Sync2(new(User), new(Point))

	u := User{
		Id:         `2`,
		Name:       `Silva`,
		Birth:      35,
		Created:    25,
		Updated_at: 231,
	}

	p := Point{
		User_id:    u.Id,
		Points:     10,
		Max_points: 10,
	}
	insertMysql(&u, &p)

}

type User struct {
	Id         string
	Name       string
	Birth      int64
	Created    int64
	Updated_at int64
}
type Point struct {
	User_id    string
	Points     int64
	Max_points int
}
