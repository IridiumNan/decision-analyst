package models

import (
	"log"
	"os"
	"time"
)

const (
	MainDBFileName = "main.db"
	EmptyStr       = ""
	NonePath       = EmptyStr
	AppName        = "decision-analyst"
)

var (
	Local    *time.Location
	userHome = NonePath
)

func init() {
	var err error
	Local, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("while setting locale: %s", err.Error())
	}
}

// HomeDir return user Home dir
// if failed, panic
func HomeDir() string {
	if userHome != NonePath {
		return userHome
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error while getting user home, err: %s", err.Error())
	}

	userHome = home

	return userHome
}
