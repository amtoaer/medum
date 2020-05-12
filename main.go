package main

import (
	"fmt"
	"log"
	"medum/config"
	"medum/output"
	"medum/public"
	"medum/sqlite"
	"medum/text"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kyokomi/emoji"
	"github.com/urfave/cli/v2"
)

var (
	begin  string
	end    string
	name   string
	remove bool
	done   int
)

func main() {
	app := &cli.App{
		Name:  "medum",
		Usage: "a terminal todo manager",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "begin",
				Aliases:     []string{"b"},
				Usage:       "event's begin time",
				Destination: &begin,
			},
			&cli.StringFlag{
				Name:        "end",
				Aliases:     []string{"e"},
				Usage:       "event's end time",
				Destination: &end,
			},
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Usage:       "event's name",
				Destination: &name,
			},
			&cli.BoolFlag{
				Name:        "remove",
				Aliases:     []string{"r"},
				Usage:       "remove outdated events",
				Destination: &remove,
			},
			&cli.IntFlag{
				Name:        "done",
				Aliases:     []string{"d"},
				Usage:       "the id of your finished event",
				Destination: &done,
			},
		},
		Action: func(c *cli.Context) error {
			conf := *config.ReadConfig()
			if remove { // if -r, ignore others
				sqlite.UpdateSqliteDB(text.MarkOutdate)
				err := sqlite.DeleteOutDate()
				if err != nil {
					fmt.Printf(text.DeleteOutDateError, err)
					os.Exit(1)
				}
				fmt.Println(text.DeleteSuccess)
			} else if done != 0 { // if -d, ignore others
				err := sqlite.DeleteID(done)
				if err != nil {
					fmt.Printf(text.DeleteIDError, err)
					os.Exit(1)
				}
			} else {
				if len(end) == 0 { //if not -e, print todo list
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
						output.Call(conf.TimeColor, "%s", formatTime(t.BeginDate, t.EndDate, t.IsEnd))
						// \n
						fmt.Println()
					}
				} else { //-e
					var beginTime, endTime time.Time
					endTime = handleString(end)
					if len(begin) == 0 {
						beginTime = time.Now()
					} else {
						beginTime = handleString(begin)
					}
					//if not -b, use time.Now(), else use -b
					if beginTime.Unix() >= endTime.Unix() {
						fmt.Println(text.TimeError)
						os.Exit(1)
					}
					err := sqlite.InsertSqliteDB(name, beginTime, endTime)
					if err != nil {
						fmt.Printf(text.InsertDBError, err)
					}
					sqlite.UpdateSqliteDB(text.MarkNotStart)
					fmt.Println(text.InsertSuccess)
				}
			}
			return nil
		}}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
func handleString(tm string) time.Time {
	tmp := strings.Split(tm, ".")
	if len(tmp) != 4 {
		fmt.Println(text.LengthError)
		os.Exit(1)
	} else {
		stdString := fmt.Sprintf(text.MyTime, strconv.Itoa(time.Now().Year()), tmp[0], tmp[1], tmp[2], tmp[3])
		result, err := time.ParseInLocation(text.StandardTime, stdString, time.Now().Local().Location())
		if err != nil {
			fmt.Printf(text.ParamError, err)
			os.Exit(1)
		}
		return result
	}
	//useless line, just to prevent warning
	return time.Now()
}

func formatTime(begin, end int64, IsEnd uint8) string {
	now := time.Now()
	if IsEnd == 0 {
		beginTime := time.Unix(begin, 0)
		dur := beginTime.Sub(now)
		return fmt.Sprintf(text.TimeBeginning, formatTimeString(dur.Minutes()))
	} else if IsEnd == 1 {
		endTime := time.Unix(end, 0)
		dur := endTime.Sub(now)
		return fmt.Sprintf(text.TimeRemaining, formatTimeString(dur.Minutes()))
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
	return tmp
}
