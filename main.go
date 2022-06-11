package main

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"fmt"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/alecthomas/kong"

	// log "github.com/sirupsen/logrus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func FetchDepartures(station *Station, limit int) (deps []Departure, err error) {
	ctx := context.Background()
	client := graphql.NewClient("https://api.entur.io/journey-planner/v2/graphql", http.DefaultClient)

	resp, err := StopNearMe(ctx, client, station.Code, limit)
	if err != nil {
		log.Error().Msg("Unable to fetch departures")
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
		log.Info().Str("code", c.Code).Str("name", c.Name).Msg("Unable to create instance")
	}
	s.Print()
	return nil
}

type RmCmd struct {
	ID uint `arg:""`
}

type AssocNetworkCmd struct {
    ID int `arg:""`
}

func (c *AssocNetworkCmd) Run() error {
	db := GetDbConnection()
	var station Station
	if err := db.Find(&station, c.ID).Error; err != nil {
		msg := fmt.Errorf("Unable to find any station with id %d.", c.ID)
		log.Error().Int("id", c.ID).Msg(msg.Error())
		return err
	}

	ssid, err := GetCurrentNetworkName()
	if err != nil {
		// log.WithField("ssid", ssid).Error("foo")
		return err
	}

	// do we already have a row containing this ssid?
	var count int64
	if db.Model(&WifiToStationBinding{}).Where("Name = ?", ssid).Count(&count); count > 0 {
		return errors.New("Already exists a binding with the same ssid")
	}

	obj := WifiToStationBinding{
		Name:    ssid,
		Station: station,
	}

	if err := db.Create(&obj).Error; err != nil {
		log.Error().Err(err)
		return err
	}
	log.Info().Msgf("OK, associated station '%s' to network with ssid '%s' ", station.Name, ssid)

	return nil
}

func GetCurrentNetworkName() (string, error) {
	out, err := exec.Command("iwgetid").Output()
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`"(.*)"`)
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
		log.Error().Msgf("Did not find any row with id %d. No rows were affected", c.ID)
		return nil
	}
	if err := x.Error; err != nil {
		log.Error().Err(err)
		return err
	}
	fmt.Println(c.ID)
	return nil
}

var CLI struct {
	Debug        bool            `help:"Enable debug mode"`
	Show         ShowCmd         `cmd:"" help:"List info for station"`
	Add          AddCmd          `cmd:"" help:"add a new station to database"`
	Rm           RmCmd           `cmd:"" help:"Remove a station by its id"`
	AssocNetwork AssocNetworkCmd `cmd:""`

	Ls LsCmd `cmd:"" help:"List stored stations"`
}

func (c *ShowCmd) Run() error {
	log.Info().Str("param:station", c.Station).Int("param:limit", c.Limit).Msg("Run started")
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
			log.Error().Msg("Found no rows with that name")
		} else if len(stops) > 1 {
			var names []string
			for _, s := range stops {
				names = append(names, s.Name)
			}

			log.Error().Msgf("Several matching stations found; %s", strings.Join(names, ", "))
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
	deps, _ := FetchDepartures(&stop, c.Limit)
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

	// ./depts add Frydenlund            NSR:StopPlace:58405
	// ./depts add Forskningsparken   NSR:StopPlace:59600
	// ./depts add Skullerud NSR:StopPlace:58227

	ctx := kong.Parse(&CLI)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	err := ctx.Run()
	if err != nil {
        // log.Info().Msg("Hello world")
		log.Error().Err(err).Msg("")
	}

}
