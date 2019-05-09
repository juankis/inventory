package dao

import (
	"database/sql"
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

//RowExists function to check if exist result from some query
func RowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := Db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Errorf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}

//CheckIfExistValue func to check if exist some value in x column and table x
func CheckIfExistValue(column string, value string, table string) bool {
	return RowExists("SELECT "+column+" FROM `"+table+"` WHERE "+column+"= ?", value)
}
