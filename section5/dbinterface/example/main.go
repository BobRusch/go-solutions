package main

import (
	"github.com/bobrusch/go-solutions/section5/database"
	"github.com/bobrusch/go-solutions/section5/dbinterface"
	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// this wont do anything if commit is successful
	defer tx.Rollback()

	if err := dbinterface.Exec(tx); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
