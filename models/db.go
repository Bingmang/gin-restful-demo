package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

const connStr = "host=202.204.62.32 port=5432 user=atp password=atp dbname=atp sslmode=disable"

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Panicln(err)
	}
}
