// Package sqlite is used to interactive with sqlite database
package sqlite

import (
	"database/sql"
	"fmt"
	"medum/path"
	"medum/text"
	"os"
	"time"

	//sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

func openSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", path.GetDataPath())
	if err != nil {
		fmt.Printf(text.OpenDBError, err)
		os.Exit(1)
	}
	db.Exec(text.CreateTable)
	return db
}

// InsertSqliteDB is used to insert event line to sqliteDB
func InsertSqliteDB(eventContent string, beginDate, endDate time.Time) error {
	db := openSqliteDB()
	defer db.Close()
	_, err := db.Exec(text.InsertRow, eventContent, beginDate.Unix(), endDate.Unix())
	return err
}

//UpdateSqliteDB is used to exec sql string.
//sql string can be text.MarkNotStart, text.MarkInProgress, text.MarkOutdate.
func UpdateSqliteDB(sql string) error {
	db := openSqliteDB()
	defer db.Close()
	_, err := db.Exec(sql, time.Now().Unix())
	return err
}

// QuerySqliteDB returns the result of query.
func QuerySqliteDB() *sql.Rows {
	db := openSqliteDB()
	defer db.Close()
	result, err := db.Query(text.QueryRow)
	if err != nil {
		fmt.Printf(text.QueryDBError, err)
		os.Exit(1)
	}
	return result
}

//DeleteOutDate removes those outdated events
func DeleteOutDate() error {
	db := openSqliteDB()
	defer db.Close()
	_, err := db.Exec(text.DeleteOutDate)
	return err
}

//DeleteID removes a line according to specific ID
func DeleteID(id int) error {
	db := openSqliteDB()
	defer db.Close()
	_, err := db.Exec(text.DeleteID, id)
	return err
}
