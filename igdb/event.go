package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct Event -add-tags json -w

// Event represents gaming events such as E3, PAX, and Gamescom.
// For more information visit: https://api-docs.igdb.com/#event
type Event struct {
	ID            int    `json:"id"`
	CreatedAt     int    `json:"created_at"`
	Description   string `json:"description"`
	EndTime       int    `json:"end_time"`
	EventLogo     int    `json:"event_logo"`
	EventNetworks []int  `json:"event_networks"`
	Games         []int  `json:"games"`
	LiveStreamURL string `json:"live_stream_url"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	StartTime     int    `json:"start_time"`
	TimeZone      string `json:"time_zone"`
	UpdatedAt     int    `json:"updated_at"`
	Videos        []int  `json:"videos"`
}

// EventService handles all the API calls for the IGDB Event endpoint.
type EventService service

// Get returns a single Event identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Events, an error is returned.
func (cs *EventService) Get(id int, opts ...Option) (*Event, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var event []*Event

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &event, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Event with ID %v", id)
	}

	return event[0], nil
}

// List returns a list of Events identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Event is ignored. If none of the IDs
// match a Event, an error is returned.
func (cs *EventService) List(ids []int, opts ...Option) ([]*Event, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var event []*Event

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &event, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Events with IDs %v", ids)
	}

	return event, nil
}

// Index returns an index of Events based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Events can
// be found using the provided options, an error is returned.
func (cs *EventService) Index(opts ...Option) ([]*Event, error) {
	var event []*Event

	err := cs.client.post(cs.end, &event, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Events")
	}

	return event, nil
}

// Search returns a list of Events found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Events are found using the provided query, an error is returned.
func (cs *EventService) Search(qry string, opts ...Option) ([]*Event, error) {
	var event []*Event

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &event, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Event with query %s", qry)
	}

	return event, nil
}

// Count returns the number of Events available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Events to count.
func (cs *EventService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Events")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Event object.
func (cs *EventService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Event fields")
	}

	return f, nil
}
