package main

import (
	"fmt"
	"medum/config"
)

func main() {
	test := *config.ReadConfig()
	fmt.Println(test)
}
