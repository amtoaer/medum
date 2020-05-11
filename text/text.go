//Package text contains all strings which may be printed
package text

//error list
const (
	HomedirError      = "error while getting homedir: %v\n"
	OpenConfigError   = "error while opening config: %v\n"
	CreateConfigError = "error while creating config: %v\n"
	OpenDBError       = "error while opening DB: %v\n"
	InsertDBError     = "error while inserting DB: %v\n"
	QueryDBError      = "error while querying DB: %v\n"
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
	QueryRow       = `select * from eventList where isEnd<>2 order by isEnd,endDate`
	DeleteOutDate  = `delect from eventList where isEnd=2`
)
