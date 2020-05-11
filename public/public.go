//Package public includes public data type like configuration
package public

// Configuration is the config struct.
type Configuration struct {
	//under consideration,these are just examples for debug
	NumberColor string
	EventColor  string
	TimeColor   string
}

// Event struct
type Event struct {
	ID           uint64
	EventContent string
	BeginDate    int64
	EndDate      int64
	IsEnd        uint8
}
