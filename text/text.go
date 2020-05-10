//Package text contains all strings which may be printed
package text

//error list
const (
	HomedirError      = "error while getting homedir: %v\n"
	OpenConfigError   = "error while opening config: %v\n"
	CreateConfigError = "error while creating config: %v\n"
	OpenDBError       = "error while opening DB: %v\n"
	InsertDBError     = "error while inserting DB: %v\n"
)

//sql list
const (
	//sqlite doesn't support bool type, so use integer.
	CreateTable = `create table if not exists eventList(
		id integer primary key autoincrement,
		event varchar(100) not null,
		beginDate datetime not null,
		endDate datetime not null,
		isEnd integer default 0 not null
	);`
	InsertTable = `insert into eventList (event,beginDate,endDate) values (?,?,?)`
)
