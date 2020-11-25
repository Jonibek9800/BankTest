package modules

import _ "database/sql"

type Users struct {
	Id int
	Name string
	Surname string
	Age  int
	Sex  string
	Login string
	Password string
	remove bool
}

