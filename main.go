package main

import (
	"context"
	"os/exec"
	"regexp"
	"strings"

	"fmt"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

func FetchDepartures(station *Station, limit int) (deps []Departure, err error) {
	ctx := context.Background()
	client := graphql.NewClient("https://api.entur.io/journey-planner/v2/graphql", http.DefaultClient)

	resp, err := StopNearMe(ctx, client, station.Code)
	if err != nil {
		log.Error("Unable to fetch departures")
		return nil, err
	}
	xaa := resp.StopPlace.EstimatedCalls

	for _, s := range xaa {
		t, err := time.Parse("2006-01-02T15:04:05-0700", s.AimedArrivalTime)
		if err != nil {
			panic(err)
		}
		dep := Departure{
			FromStation:   station,
			PublicCode:    s.ServiceJourney.Line.PublicCode,
			Display:       s.DestinationDisplay.FrontText,
			DepartureTime: t,
		}
		deps = append(deps, dep)
	}
	return deps, nil
}

type Context struct {
	Debug bool `help:"Enable debug mode" default:"false"`
}

type ShowCmd struct {
	Station string `arg:"" help:"station name. If STATION exactly matches a station name, it is used. Otherwise, if the STATION exactly matches the prefix of a station, that station will be used."`
	Limit   int    `help:"Number of stops to display" default:"5"`
}

type LsCmd struct {
}

type AddCmd struct {
	Name string `arg:"" help:"station name"`
	Code string `arg:"" help:"station code"`
}

func (c *AddCmd) Run() error {
	station := Station{Code: c.Code, Name: c.Name}
	var s Station
	if err := GetDbConnection().FirstOrCreate(&s, &station).Error; err != nil {
		log.WithField("code", c.Code).WithField("name", c.Name).Fatalf("Unable to create instance: %v")
	}
	s.Print()
	return nil
}

type RmCmd struct {
	ID uint `arg:""`
}

type TestCmd struct {
}

func GetCurrentNetworkName() (string, error) {
	out, err := exec.Command("iwgetid").Output()
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`"(xxx.*)"`)
	match := re.FindSubmatch(out)
	if len(match) == 0 {
		return "", fmt.Errorf("Unable to find name from output '%s'", out)
	}
	capture := match[1]
	return string(capture), nil
}

func (c *RmCmd) Run() error {
	db := GetDbConnection()
	var s Station
	x := db.Delete(&s, c.ID)
	if x.RowsAffected == 0 {
		log.Errorf("Did not find any row with id %d. No rows were affected", c.ID)
		return nil
	}
	if err := x.Error; err != nil {
		log.Error(err)
		return err
	}
	fmt.Println(c.ID)
	return nil
}

var CLI struct {
	Debug bool    `help:"Enable debug mode"`
	Show  ShowCmd `cmd:"" help:"List info for station"`
	Add   AddCmd  `cmd:"" help:"add a new station to database"`
	Rm    RmCmd   `cmd:"" help:"Remove a station by its id"`
	Test  TestCmd `cmd:""`

	Ls LsCmd `cmd:"" help:"List stored stations"`
}

func (c *ShowCmd) Run() error {
	log.WithFields(log.Fields{"param:station": c.Station, "param:limit": c.Limit}).Info("Run started")
	var stop Station
	// var q *gorm.DB
	db := GetDbConnection()
	// TODO: do not fail if we don't find it with exact match...
	var count int64
	db.Where("name = ?", c.Station).Count(&count)
	if count == 0 {
		// let's try inexact...
		var stops []Station
		db.Where("lower(name) like ?", c.Station+"%").Find(&stops)
		if len(stops) == 0 {
			log.Error("Found no rows with that name")
		} else if len(stops) > 1 {
			var names []string
			for _, s := range stops {
				names = append(names, s.Name)
			}

			log.Errorf("Several matching stations found; %s", strings.Join(names, ", "))
		} else {
			stop = stops[0]
		}
		// log.Errorf("AAA %d" , q.RowsAffected)
	} else {
		db.Where("name = ?", c.Station).First(&stop)
	}

	// if err := GetDbConnection().Where("name = ?", c.Station).First(&stop).Error; err != nil {
	// 	log.Error("Got error when fetching station")
	// 	// return
	// 	os.Exit(1)
	// }
	deps, _ := FetchDepartures(&stop, 5)
	for _, s := range deps {
		s.Print()
	}
	return nil
}

func (c *LsCmd) Run() error {
	db := GetDbConnection()
	var stations []Station
	db.Find(&stations)
	for _, s := range stations {
		s.Print()
	}
	return nil
}

func main() {

	// log.SetFormatter(&log.JSONFormatter{PrettyPrint: false})
	// log.SetLevel(log.ErrorLevel)
	// Frydenlund            NSR:StopPlace:58405
	// Forskningsparken   NSR:StopPlace:59600
	// Skullerud NSR:StopPlace:58227

	ctx := kong.Parse(&CLI)

	if CLI.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	ctx.Run()
}
