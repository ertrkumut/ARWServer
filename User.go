package main

import "strconv"

type User struct {
	name    string
	id      int
	session Session
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) SetSession(sess Session) {
	u.session = sess
}

func (u *User) GetDataForOtherUser(user User) string {
	userData := u.name + "^^" + strconv.Itoa(u.id) + "^^"
	if u != &user {
		userData += "false" // IsMe true ^^ false
	} else {
		userData += "true"
	}

	return userData
}
