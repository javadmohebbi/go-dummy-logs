package main

import (
	"fmt"

	godummylogs "github.com/javadmohebbi/go-dummy-logs"
)

// config gile
var conf_file = "/Users/javad/Projects/go-dummy-logs/godummylogs.example.conf"

func main() {
	dummy := godummylogs.New(conf_file)
	fmt.Println(dummy.Config)
}
