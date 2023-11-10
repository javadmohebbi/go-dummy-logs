package godummylogs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Users struct {
	Username string `json:"username"`
	UserID   uint   `json:"userID"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type Config struct {
	Hostnames  []string `json:"hostnames"`
	KnownUsers []Users  `json:"knownUsers"`

	ContextPath string `json:"context_path"`

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
