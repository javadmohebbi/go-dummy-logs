package main

import (
	godummylogs "github.com/javadmohebbi/go-dummy-logs"
)

// Path to config.json file
var conf_file = "/Users/javad/Projects/go-dummy-logs/godummylogs.example.json"

func main() {
	dummy := godummylogs.New(conf_file)
	// c := dummy.Config
	// j, _ := json.MarshalIndent(c, "", " ")
	// fmt.Println(string(j))
	dummy.Generate()
}
