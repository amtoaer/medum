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
	tmp, _ := db.Prepare(text.InsertTable)
	_, err := tmp.Exec(eventContent, beginDate, endDate)
	return err
}
