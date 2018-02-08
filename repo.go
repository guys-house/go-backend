package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var mydb *sql.DB

func init(){
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/mydb")
	CheckErr(err)
	mydb = db
}

func CreateRecord(record Record) error{
		_, err := mydb.Exec(`INSERT INTO cash_balance (item, amount, start_date,
			end_date, update_at) values (?,?,?,?,?)`,
			record.Item, record.Amount, record.Start_date, record.End_date, time.Now())
		return err
}

func RemoveRecord(id int) error {
	_, err := mydb.Exec("DELETE from cash_balance where id = ?", id)
	return err
}