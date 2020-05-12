//Package text contains all strings which may be printed
package text

//error list
const (
	HomedirError       = "error while getting homedir: %v"
	OpenConfigError    = "error while opening config: %v"
	CreateConfigError  = "error while creating config: %v"
	OpenDBError        = "error while opening DB: %v"
	InsertDBError      = "error while inserting DB: %v"
	QueryDBError       = "error while querying DB: %v"
	LengthError        = "error while parsing parameters: time format must be month.day.hour.minute"
	ParamError         = "error while parsing time: %v"
	TimeError          = "error while comparing time: beginTime is larger than endTime"
	DeleteOutDateError = "error while deleting outdated events: %v"
	DeleteIDError      = "error while deleting event with specific ID: %v"
)

//log list
const (
	DeleteSuccess = "successfully delete outdated events"
	InsertSuccess = "successfully insert event"
	StandardTime  = "2006-1-2 15:04:05 +0800 CST"
	MyTime        = "%s-%s-%s %s:%s:00 +0800 CST"
	TimeRemaining = "%s remaining"
	TimeBeginning = "%s beginning"
)

//sql list
const (
	//sqlite doesn't support bool/date type, so use integer.
	//isEnd can be 0,1,2
	CreateTable = `create table if not exists eventList(
		id integer primary key autoincrement,
		event varchar(100) not null,
		beginDate unsigned bigint not null,
		endDate unsigned bigint not null,
		isEnd smallint default 1 not null
	);`
	InsertRow      = `insert into eventList (event,beginDate,endDate) values (?,?,?);`
	MarkOutdate    = `update eventList set isEnd=2 where ?>endDate;`
	MarkNotStart   = `update eventList set isEnd=0 where ?<beginDate;`
	MarkInProgress = `update eventList set isEnd=1 where ? between beginDate and endDate;`
	QueryRow       = `select * from eventList order by isEnd,endDate`
	DeleteOutDate  = `delect from eventList where isEnd=2`
	DeleteID       = `delete from eventList where id=?`
)
