package godummylogs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Known struct {
	Users []Users `json:"users"`
}

type Users struct {
	Username string `json:"username"`
	UserID   uint   `json:"userID"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type Generate struct {
	NickName  string     `json:"nickName"`
	FileName  string     `json:"fileName"`
	Templates []Template `json:"templates"`
}

type Template struct {
	Interval string   `json:"interval"`
	LogLines []string `json:"logLines"`
}

type Config struct {
	Hostnames []string `json:"hostnames"`
	Known     Known    `json:"known"`

	ContextPath string `json:"context_path"`

	Generate []Generate `json:"generate"`

	Auth ConfAuthLog `json:"auth"`
}

// Read .json configuration file and parse it
func ParseConfig(c string) Config {
	b, err := os.ReadFile(c)
	if err != nil {
		panic(fmt.Sprintf("[ParseConfig:ReadFile] %v", err))
	}

	var conf Config

	err = json.Unmarshal(b, &conf)
	if err != nil {
		panic(fmt.Sprintf("[ParseConfig:Unmarshal] %v", err))
	}

	return conf
}
