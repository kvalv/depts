package main

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Station struct {
	gorm.Model

	Code string `gorm:"unique"`
	Name string
}

type WifiToStationBinding struct {
	gorm.Model
	Name      string  `gorm:"unique"`
	Station   Station `gorm:"OnDelete:CASCADE"`
	StationID uint
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
		log.Info().Msg("Reusing existing database instance")
		return db
	}
	log.Debug().Msg("Connecting to database")
	config := gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	db, err := gorm.Open(sqlite.Open("test.db"), &config)
	if err != nil {
		log.Panic().Msg("failed to connect database")
	}
	log.Debug().Msg("created db connection")
	log.Debug().Msg("applying automigration")
	err = db.AutoMigrate(&Station{}, &WifiToStationBinding{})
	if err != nil {
		log.Panic().Err(err).Msg("")
	}
	return db
}
