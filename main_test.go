package main

import (
	"APS_Project/database"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	db, err := gorm.Open("mysql", "aps:password@/apsdb_test?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("Could not connect to DB: ", err)
	} else {
		defer db.Close()
	}

	database.GlobalDB = db

	db.LogMode(true)

	log.Print("Created test DB connection!")

	ret := m.Run()

	os.Exit(ret)
}
