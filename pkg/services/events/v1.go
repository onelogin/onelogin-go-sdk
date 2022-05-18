package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errEventsV1Context = "events v1 service"

type V1Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v1.
func New(repo services.Repository, host string) *V1Service {
	return &V1Service{
		Endpoint:     fmt.Sprintf("%s/api/1/events", host),
		Repository:   repo,
		ErrorContext: errEventsV1Context,
	}
}

// Query retrieves all the events from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all events
func (svc *V1Service) Query(query *EventQuery) ([]Event, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var events []Event
	for _, bytes := range resp {
		var unmarshalled []Event
		json.Unmarshal(bytes, &unmarshalled)
		events = append(events, unmarshalled...)
	}
	return events, nil
}

// GetOne retrieves the event by id and returns it
func (svc *V1Service) GetOne(id int32) (*Event, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}
	var event Event

	if len(resp) < 1 {
		return nil, errors.New("invalid length of response returned")
	}

	json.Unmarshal(resp[0], &event)
	return &event, nil
}

// GetTypes returns a list of all OneLogin EventType available. It provides
// event type names, event type IDs, and descriptions.
func (svc *V1Service) GetTypes() ([]EventType, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s", svc.Endpoint, "types"),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}

	var eventTypes []EventType
	for _, bytes := range resp {
		var unmarshalled []EventType
		json.Unmarshal(bytes, &unmarshalled)
		eventTypes = append(eventTypes, unmarshalled...)
	}
	return eventTypes, nil
}
