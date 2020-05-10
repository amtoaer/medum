package main

import (
	"fmt"
	"medum/config"
	"medum/sqlite"
	"medum/text"
	"os"
	"time"
)

func main() {
	//try to read config file.
	conf := *config.ReadConfig()
	fmt.Println(conf)
	begin := time.Now()
	end := time.Now()
	//try to insert an event to the db.
	err := sqlite.InsertSqliteDB("test", begin, end)
	if err != nil {
		fmt.Printf(text.InsertDBError, err)
		os.Exit(1)
	}
}
