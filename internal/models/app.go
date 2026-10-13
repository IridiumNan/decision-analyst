package models

import (
	"log"
	"time"
)

var Local *time.Location

func init() {
	var err error
	Local, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("while setting locale: %s", err.Error())
	}
}

const MainDBFileName = "main.db"
