package dao

import (
	"fmt"

	//_ this library is necesary to connet with database
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	//_ this library is necesary to connet with database test
	"github.com/juankis/inventory/api/config"
)

//Db is connection with de data base
var Db = Database()

// Database middleware
func Database() *sqlx.DB {
	// Create connection pool
	dsn := fmt.Sprintf(
		config.DB_DSN_TEMPLATE,
		"root",
		"mysql123",
		"localhost:3306",
	)
	var db *sqlx.DB
	var err error
	fmt.Println(fmt.Sprintf(`[event:db_connecting][dsn:%s@%s]`, "root", "localhost"))
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(fmt.Sprintf(`[event:db_connection_error][error:"%s"]`, err.Error()))
	}
	return db
}
