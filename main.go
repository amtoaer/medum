package main

import (
	"fmt"
	"medum/config"
	"medum/public"
	"medum/sqlite"
	"medum/text"
	"os"
	"time"
)

func formatTime(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}
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
	//try to query
	test := sqlite.QuerySqliteDB()
	var t public.Event
	for test.Next() {
		test.Scan(&t.ID, &t.EventContent, &t.BeginDate, &t.EndDate, &t.IsEnd)
		fmt.Println(t.ID, t.EventContent, formatTime(t.BeginDate), formatTime(t.EndDate))
	}
}
