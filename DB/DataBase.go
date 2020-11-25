package DB

import (
	"database/sql"
)
func DBinit(db *sql.DB) (err error) {
	const usersDDL = `create table Users
(
       Id integer primary key autoincrement,
       Name text not null,
       Surename text not null,
       Age integer not null,
       Sex text not null,
       Login text not null,
       Password text not null,
       Remove boolean not null default false
      )
`
	_, err = db.Exec(usersDDL)
	if err != nil {
		return err
	}
	return nil
}

func DBcur(db *sql.DB) (err error)  {
	const currencyDDL = `create table currency
	(
		id integer primary key autoincrement,
		name text not null 
	)
`
	_, err1 := db.Exec(currencyDDL)
	if err1 != nil {
		return err1
	}
	return nil
}

func DBacount(db *sql.DB) (err error)  {
	const acountDDL = `create table account(
    Id integer not null,
    UsersId  integer ,
    Number integer not null,
    Amount integer not null,
    Currency integer not null
)
`
	_, err2 := db.Exec(acountDDL)
	if err2 != nil {
		return err2
	}
	return nil
}
