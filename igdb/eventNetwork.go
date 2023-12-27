package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct EventNetwork -add-tags json -w

// EventNetwork represents urls related to an event like twitch, youtube, etc.
// For more information visit: https://api-docs.igdb.com/#event-network
type EventNetwork struct {
	ID          int    `json:"id"`
	CreatedAt   int    `json:"created_at"`
	Event       int    `json:"event"`
	NetworkType int    `json:"network_type"`
	UpdatedAt   int    `json:"updated_at"`
	URL         string `json:"url"`
}

// EventNetworkService handles all the API calls for the IGDB EventNetwork endpoint.
type EventNetworkService service

// Get returns a single EventNetwork identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any EventNetworks, an error is returned.
func (cs *EventNetworkService) Get(id int, opts ...Option) (*EventNetwork, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var net []*EventNetwork

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &net, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get EventNetwork with ID %v", id)
	}

	return net[0], nil
}

// List returns a list of EventNetworks identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a EventNetwork is ignored. If none of the IDs
// match a EventNetwork, an error is returned.
func (cs *EventNetworkService) List(ids []int, opts ...Option) ([]*EventNetwork, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var net []*EventNetwork

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &net, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get EventNetworks with IDs %v", ids)
	}

	return net, nil
}

// Index returns an index of EventNetworks based solely on the provided functional
// options used to sort, filter, and paginate the results. If no EventNetworks can
// be found using the provided options, an error is returned.
func (cs *EventNetworkService) Index(opts ...Option) ([]*EventNetwork, error) {
	var net []*EventNetwork

	err := cs.client.post(cs.end, &net, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of EventNetworks")
	}

	return net, nil
}

// Search returns a list of EventNetworks found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no EventNetworks are found using the provided query, an error is returned.
func (cs *EventNetworkService) Search(qry string, opts ...Option) ([]*EventNetwork, error) {
	var net []*EventNetwork

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &net, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get EventNetwork with query %s", qry)
	}

	return net, nil
}

// Count returns the number of EventNetworks available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which EventNetworks to count.
func (cs *EventNetworkService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count EventNetworks")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB EventNetwork object.
func (cs *EventNetworkService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get EventNetwork fields")
	}

	return f, nil
}
