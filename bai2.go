package main

//ex2 updateBirth
func updateBirth(id string, birth int64) error {

	session := engine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		session.Rollback()
		// if returned then will rollback automatically
		//panic(err)
		return err
	}
	var user User
	_, err := session.Where("id = ?", id).Get(&user)
	if err != nil {
		session.Rollback()
		return err
	}
	u3 := User{
		Name:  user.Name + string("  updated"),
		Birth: birth,
	}
	if _, err := session.Where("id = ?", id).Table("user").Update(&u3); err != nil {
		session.Rollback()
		return err
	}
	var newpoint Point
	_, err = session.Where("user_id = ?", id).Get(&newpoint)
	if err != nil {
		session.Rollback()
		return err
	}

	if _, err := session.Where("user_id = ?", id).Table("point").Update(&Point{
		Points:     newpoint.Points + int64(10),
		Max_points: newpoint.Max_points + int(10),
	}); err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}
