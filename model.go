package main

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/rs/zerolog/log"
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

func InitializeDatabase(url string) error {
	log.Debug().Msg("Connecting to database")
	config := gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	ldb, err := gorm.Open(sqlite.Open(url), &config)
	if err != nil {
		return err
	}
	log.Debug().Msg("created db connection")
	log.Debug().Msg("applying automigration")
	err = ldb.AutoMigrate(&Station{}, &WifiToStationBinding{})
	if err != nil {
		return err
	}
	db = ldb
	return nil
}

// Returns database connection. Panics if the database is not set up
func GetDbConnection() *gorm.DB {
	// we'll use a local variable to store the database connection so it's fast to access
	// it on the second time
	if db == nil {
		log.Fatal().Msg("Database is not initialized. Please call InitializeDatabase first.")
	}
	return db
}
