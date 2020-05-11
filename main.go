package main

import (
	"fmt"
	"medum/config"
	"medum/output"
	"medum/public"
	"medum/sqlite"
	"medum/text"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/kyokomi/emoji"
)

var (
	begin  = kingpin.Flag("begin time", "event's begin time.").Short('b').String()
	end    = kingpin.Flag("end time", "event's end time.").Short('e').String()
	name   = kingpin.Flag("name", "event's name.").Short('n').String()
	delete = kingpin.Flag("delete", "delete outdated events.").Short('d').Bool()
)

func main() {
	kingpin.Parse()
	conf := *config.ReadConfig()
	if *delete { // if -d, ignore others
		sqlite.UpdateSqliteDB(text.MarkOutdate)
		sqlite.DeleteSqliteDB()
		fmt.Println(text.DeleteSuccess)
	} else {
		if len(*end) == 0 {
			sqlite.UpdateSqliteDB(text.MarkInProgress)
			sqlite.UpdateSqliteDB(text.MarkOutdate)
			test := sqlite.QuerySqliteDB()
			var t public.Event
			for test.Next() {
				test.Scan(&t.ID, &t.EventContent, &t.BeginDate, &t.EndDate, &t.IsEnd)
				output.Call(conf.NumberColor, "%s", strconv.Itoa(t.ID))
				emoji.Printf(" | :beer:")
				output.Call(conf.EventColor, "%s", t.EventContent)
				emoji.Printf(" | :hourglass:")
				output.Call(conf.TimeColor, "%s", formatTime(t.BeginDate, t.EndDate))
				// \n
				fmt.Println()
			}
		} else {
			var beginTime, endTime time.Time
			endTime = handleString(*end)
			if len(*begin) == 0 {
				beginTime = time.Now()
			} else {
				beginTime = handleString(*begin)
			}
			if beginTime.Unix() >= endTime.Unix() {
				fmt.Println(text.TimeError)
				os.Exit(1)
			}
			err := sqlite.InsertSqliteDB(*name, beginTime, endTime)
			if err != nil {
				fmt.Printf(text.InsertDBError, err)
			}
			fmt.Println(text.InsertSuccess)
		}
	}
}
func handleString(tm string) time.Time {
	tmp := strings.Split(tm, ".")
	if len(tmp) != 4 {
		fmt.Println(text.LengthError)
		os.Exit(1)
	} else {
		stdString := fmt.Sprintf(text.MyTime, strconv.Itoa(time.Now().Year()), tmp[0], tmp[1], tmp[2], tmp[3])
		result, err := time.ParseInLocation(text.StandredTime, stdString, time.Now().Local().Location())
		if err != nil {
			fmt.Printf(text.ParamError, err)
			os.Exit(1)
		}
		return result
	}
	//useless line, just to prevent warning
	return time.Now()
}

func formatTime(begin, end int64) string {
	beginTime := time.Unix(begin, 0)
	endTime := time.Unix(end, 0)
	now := time.Now()
	if now.Unix() <= begin {
		dur := beginTime.Sub(now)
		return formatTimeString(dur.Minutes())
	} else if now.Unix() <= end {
		dur := endTime.Sub(now)
		return formatTimeString(dur.Minutes())
	} else {
		return fmt.Sprintf(text.TimeRemaining, "no time")
	}
}

func formatTimeString(min float64) string {
	var tmp string
	if min > 1440 {
		tmp = strconv.Itoa(int(min/1440)) + " days"
	} else if min > 60 {
		tmp = strconv.Itoa(int((min / 60))) + " hours"
	} else {
		tmp = strconv.Itoa(int(min)) + " minutes"
	}
	return fmt.Sprintf(text.TimeRemaining, tmp)
}
