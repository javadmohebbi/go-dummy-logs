package main

import (
	"fmt"

	godummylogs "github.com/javadmohebbi/go-dummy-logs"
)

// Path to config.json file
var conf_file = "/Users/javad/Projects/go-dummy-logs/godummylogs.example.json"

func main() {
	dummy := godummylogs.New(conf_file)
	fmt.Println(dummy.Config)
}
