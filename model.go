package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Station struct {
	gorm.Model

	Code string `gorm:"unique"`
	Name string
}

func (s *Station) Print() {
	fmt.Printf("%3d   %-20s %20s\n", s.ID, s.Name, s.Code)
}

type Departure struct {
	FromStation   *Station
	PublicCode    string
	Display       string
	DepartureTime time.Time
}

func (s *Departure) Print() {
	mins := s.DepartureTime.Sub(time.Now()).Minutes()
	fmt.Printf("%-3s %-25s %5.1fm\n", s.PublicCode, s.Display, mins)
}

var db *gorm.DB

func GetDbConnection() *gorm.DB {
	// we'll use a local variable to store the database connection so it's fast to access
	// it on the second time
	if db != nil {
		log.Info("Reusing existing database instance")
		return db
	}
	log.Info("Connecting to database")
	config := gorm.Config{}
	db, err := gorm.Open(sqlite.Open("test.db"), &config)
	if err != nil {
		log.Panic("failed to connect database")
	}
	log.Info("created db connection")
	log.Info("applying automigration")
	err = db.AutoMigrate(&Station{})
	if err != nil {
		log.Panic(err)
	}
	return db
}

func AddStation(code, name string) {
	db := GetDbConnection()
	var s Station
	res := db.FirstOrCreate(&s, &Station{Code: code, Name: name})
	if res.Error != nil {
		log.Panic(res)
	}
}

func GetFrydenlund() Station {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db := GetDbConnection()
	// if err != nil {
	// 	panic(err)
	// }
	var o Station
	db.First(&o)
	return o
}

type LocalNetworkBinding struct {
}

var favourites = []Station{
	{Name: "Frydenlund", Code: "NSR:StopPlace:58405"},
	{Name: "Forskningsparken", Code: "NSR:StopPlace:59600"},
}
