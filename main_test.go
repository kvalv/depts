package main

import (
	"testing"
)

var Frydenlund Station = Station{Code: "NSR:StopPlace:58405", Name: "Frydenlund"}
var Skullerud Station = Station{Code: "NSR:StopPlace:58227", Name: "Skullerud"}
var Forskningsparken Station = Station{Code: "NSR:StopPlace:59600", Name: "Forskningsparken"}

func setup() {
	InitializeDatabase(":memory:")
	db := GetDbConnection()
	db.Create(&Frydenlund)
	db.Create(&Skullerud)
	db.Create(&Forskningsparken)
}
func teardown() {
}

func TestResolveStationName(t *testing.T) {
	setup()
	defer teardown()

	testCases := []struct {
		input     string
		want      string
		shouldErr bool
		descr     string
	}{
		{"Skullerud", "Skullerud", false, "full name"},
		{"fr", "Frydenlund", false, "prefix name"},
		{"f", "", true, "ambiguous name"},  // matches 'frydenlund' and 'forskningsparken'
		{"xxx", "", true, "no match"},
	}

	for _, tc := range testCases {
		t.Run(tc.descr, func(t *testing.T) {
			stop, err := resolveStationName(tc.input)
			if tc.shouldErr && err == nil {
				t.Fatalf("test case should give an error but did not.")
			}
			if !tc.shouldErr && err != nil {
				t.Fatalf("test case should not give an error, but gave the error '%s'", err.Error())
			}
			if stop.Name != tc.want {
				t.Fatalf("got station with name '%s' but wanted station with name '%s'", stop.Name, tc.want)
			}
		})
	}

}
