// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// StopNearMeResponse is returned by StopNearMe on success.
type StopNearMeResponse struct {
	// Get a single stopPlace based on its id)
	StopPlace StopNearMeStopPlace `json:"stopPlace"`
}

// GetStopPlace returns StopNearMeResponse.StopPlace, and is useful for accessing the field via an interface.
func (v *StopNearMeResponse) GetStopPlace() StopNearMeStopPlace { return v.StopPlace }

// StopNearMeStopPlace includes the requested fields of the GraphQL type StopPlace.
// The GraphQL type's documentation follows.
//
// Named place where public transport may be accessed. May be a building complex (e.g. a station) or an on-street location.
type StopNearMeStopPlace struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Whether this stop place is suitable for wheelchair boarding.
	WheelchairBoarding WheelchairBoarding `json:"wheelchairBoarding"`
	// The transport mode serviced by this stop place.
	TransportMode TransportMode `json:"transportMode"`
	// List of visits to this stop place as part of vehicle journeys.
	EstimatedCalls []StopNearMeStopPlaceEstimatedCallsEstimatedCall `json:"estimatedCalls"`
}

// GetId returns StopNearMeStopPlace.Id, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlace) GetId() string { return v.Id }

// GetName returns StopNearMeStopPlace.Name, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlace) GetName() string { return v.Name }

// GetDescription returns StopNearMeStopPlace.Description, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlace) GetDescription() string { return v.Description }

// GetWheelchairBoarding returns StopNearMeStopPlace.WheelchairBoarding, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlace) GetWheelchairBoarding() WheelchairBoarding { return v.WheelchairBoarding }

// GetTransportMode returns StopNearMeStopPlace.TransportMode, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlace) GetTransportMode() TransportMode { return v.TransportMode }

// GetEstimatedCalls returns StopNearMeStopPlace.EstimatedCalls, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlace) GetEstimatedCalls() []StopNearMeStopPlaceEstimatedCallsEstimatedCall {
	return v.EstimatedCalls
}

// StopNearMeStopPlaceEstimatedCallsEstimatedCall includes the requested fields of the GraphQL type EstimatedCall.
// The GraphQL type's documentation follows.
//
// List of visits to quays as part of vehicle journeys. Updated with real time information where available
type StopNearMeStopPlaceEstimatedCallsEstimatedCall struct {
	// Whether this call has been updated with real time information.
	Realtime bool `json:"realtime"`
	// Scheduled time of arrival at quay. Not affected by read time updated
	AimedArrivalTime string `json:"aimedArrivalTime"`
	// Whether vehicle may be boarded at quay.
	ForBoarding        bool                                                             `json:"forBoarding"`
	DestinationDisplay StopNearMeStopPlaceEstimatedCallsEstimatedCallDestinationDisplay `json:"destinationDisplay"`
	ServiceJourney     StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourney     `json:"serviceJourney"`
}

// GetRealtime returns StopNearMeStopPlaceEstimatedCallsEstimatedCall.Realtime, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCall) GetRealtime() bool { return v.Realtime }

// GetAimedArrivalTime returns StopNearMeStopPlaceEstimatedCallsEstimatedCall.AimedArrivalTime, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCall) GetAimedArrivalTime() string {
	return v.AimedArrivalTime
}

// GetForBoarding returns StopNearMeStopPlaceEstimatedCallsEstimatedCall.ForBoarding, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCall) GetForBoarding() bool { return v.ForBoarding }

// GetDestinationDisplay returns StopNearMeStopPlaceEstimatedCallsEstimatedCall.DestinationDisplay, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCall) GetDestinationDisplay() StopNearMeStopPlaceEstimatedCallsEstimatedCallDestinationDisplay {
	return v.DestinationDisplay
}

// GetServiceJourney returns StopNearMeStopPlaceEstimatedCallsEstimatedCall.ServiceJourney, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCall) GetServiceJourney() StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourney {
	return v.ServiceJourney
}

// StopNearMeStopPlaceEstimatedCallsEstimatedCallDestinationDisplay includes the requested fields of the GraphQL type DestinationDisplay.
// The GraphQL type's documentation follows.
//
// An advertised destination of a specific journey pattern, usually displayed on a head sign or at other on-board locations.
type StopNearMeStopPlaceEstimatedCallsEstimatedCallDestinationDisplay struct {
	// Name of destination to show on front of vehicle.
	FrontText string `json:"frontText"`
}

// GetFrontText returns StopNearMeStopPlaceEstimatedCallsEstimatedCallDestinationDisplay.FrontText, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCallDestinationDisplay) GetFrontText() string {
	return v.FrontText
}

// StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourney includes the requested fields of the GraphQL type ServiceJourney.
// The GraphQL type's documentation follows.
//
// A planned vehicle journey with passengers.
type StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourney struct {
	Id   string                                                           `json:"id"`
	Line StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine `json:"line"`
}

// GetId returns StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourney.Id, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourney) GetId() string { return v.Id }

// GetLine returns StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourney.Line, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourney) GetLine() StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine {
	return v.Line
}

// StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine includes the requested fields of the GraphQL type Line.
// The GraphQL type's documentation follows.
//
// A group of routes which is generally known to the public by a similar name or number
type StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine struct {
	Name          string        `json:"name"`
	TransportMode TransportMode `json:"transportMode"`
	// Publicly announced code for line, differentiating it from other lines for the same operator.
	PublicCode string `json:"publicCode"`
}

// GetName returns StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine.Name, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine) GetName() string {
	return v.Name
}

// GetTransportMode returns StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine.TransportMode, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine) GetTransportMode() TransportMode {
	return v.TransportMode
}

// GetPublicCode returns StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine.PublicCode, and is useful for accessing the field via an interface.
func (v *StopNearMeStopPlaceEstimatedCallsEstimatedCallServiceJourneyLine) GetPublicCode() string {
	return v.PublicCode
}

type TransportMode string

const (
	TransportModeAir       TransportMode = "air"
	TransportModeBus       TransportMode = "bus"
	TransportModeCableway  TransportMode = "cableway"
	TransportModeWater     TransportMode = "water"
	TransportModeFunicular TransportMode = "funicular"
	TransportModeLift      TransportMode = "lift"
	TransportModeRail      TransportMode = "rail"
	TransportModeMetro     TransportMode = "metro"
	TransportModeTram      TransportMode = "tram"
	TransportModeCoach     TransportMode = "coach"
	TransportModeUnknown   TransportMode = "unknown"
)

type WheelchairBoarding string

const (
	// There is no accessibility information for the stopPlace/quay.
	WheelchairBoardingNoinformation WheelchairBoarding = "noInformation"
	// Boarding wheelchair-accessible serviceJourneys is possible at this stopPlace/quay.
	WheelchairBoardingPossible WheelchairBoarding = "possible"
	// Wheelchair boarding/alighting is not possible at this stop.
	WheelchairBoardingNotpossible WheelchairBoarding = "notPossible"
)

// __StopNearMeInput is used internally by genqlient
type __StopNearMeInput struct {
	Id string `json:"id"`
}

// GetId returns __StopNearMeInput.Id, and is useful for accessing the field via an interface.
func (v *__StopNearMeInput) GetId() string { return v.Id }

func StopNearMe(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*StopNearMeResponse, error) {
	__input := __StopNearMeInput{
		Id: id,
	}
	var err error

	var retval StopNearMeResponse
	err = client.MakeRequest(
		ctx,
		"StopNearMe",
		`
query StopNearMe ($id: String!) {
	stopPlace(id: $id) {
		id
		name
		description
		wheelchairBoarding
		transportMode
		estimatedCalls(numberOfDepartures: 5) {
			realtime
			aimedArrivalTime
			forBoarding
			destinationDisplay {
				frontText
			}
			serviceJourney {
				id
				line {
					name
					transportMode
					publicCode
				}
			}
		}
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}
