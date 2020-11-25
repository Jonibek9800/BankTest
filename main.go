package main

import (
	"HumoLab/DBwork/DB"
	"HumoLab/DBwork/modules"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// first struct
	db, err := sql.Open("sqlite3", "Bank")
	if err != nil {
		log.Fatal("Data Base conection is failed", err)
	}
	err = DB.DBinit(db)
	if err != nil {
		fmt.Println("Sorry not found", err)
	}
	//second struct
	open, err1 := sql.Open("sqlite3", "Bank")
	if err1 != nil {
		log.Fatal("sorry not found", err1)
	}
	err1 = DB.DBcur(open)
	if err1 != nil {
		fmt.Println("404 not found", err1, )
	}

	//third struct
	sd, err2 := sql.Open("sqlite3", "Bank")
	if err2 != nil {
		log.Fatal("sorry mistake address", err2)
	}
	err2 = DB.DBacount(sd)
	if err2 != nil {
		fmt.Println(" :), ;-* sory", err2)

	}
	acount := modules.Account{
		Id:       1,
		UsersId:  1,
		Number:   105,
		Amount:   1000,
		Currency: "tjs",
	}
	user := modules.Users{
		Id:       10,
		Name:     "Jonibek",
		Surname:  "Mahmudov",
		Age:      22,
		Sex:      "man",
		Login:    "25248",
		Password: "1588899",
	}
	err3 := AddNewClient(db, user)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(user)

	err4 := AddNewAcount(db, acount)
	if err4 != nil {
		fmt.Println(err4)
	}
		fmt.Println(acount)

		err5 := RemoveById(db)
		if err5 != nil {
			fmt.Println(err5)
		}
		fmt.Println("Sorry I can not remove")

	err6 := Updute(db)
	if err6 != nil {
		fmt.Println(err6)
	}
	fmt.Println("Sorry nothing update")
}
func AddNewClient(db *sql.DB, client modules.Users) (err3 error) {

	_, err3 = db.Exec("Insert Into Clients(Name, Surname, Age, Sex, Login, Password) values (($1), ($2), ($3), ($4), ($5), ($6))", client.Name, client.Surname, client.Age, client.Sex, client.Login, client.Password)
	if err3 != nil {
		fmt.Println(err3)
	}
	return err3
}

func AddNewAcount(db *sql.DB, acount modules.Account) (err4 error)  {
	_, err4 = db.Exec("Insert Into Account(UsersId, Number, Amount, Currency) values (($1), ($2), ($3), ($4))", acount.UsersId, acount.Number, acount.Amount, acount.Currency)
	if err4 != nil {
		fmt.Println(err4)
	}
	return err4
}


func RemoveById(db *sql.DB) (err5 error)  {
	_, err5 = db.Exec("Delete From Users Where Id = 1")
	if err5 != nil {
		fmt.Println(err5)
	}
	return err5
}

func Updute(db *sql.DB) (err6 error) {
	_, err6 = db.Exec("Update Account Set Amount = Amount + 10000 ")
	if err6 != nil {
		fmt.Println(err6)
	}
	return err6
}
//Delete and Update сделани с помощью консоли