package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct NetworkType -add-tags json -w

// NetworkType represents social network related to event such as Facebook or Twitter.
// For more information visit: https://api-docs.igdb.com/#network-type
type NetworkType struct {
	ID            int    `json:"id"`
	CreatedAt     int    `json:"created_at"`
	EventNetworks []int  `json:"event_networks"`
	Name          string `json:"name"`
	UpdatedAt     int    `json:"updated_at"`
}

// NetworkTypeService handles all the API calls for the IGDB NetworkType endpoint.
type NetworkTypeService service

// Get returns a single NetworkType identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any NetworkTypes, an error is returned.
func (cs *NetworkTypeService) Get(id int, opts ...Option) (*NetworkType, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var net []*NetworkType

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &net, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get NetworkType with ID %v", id)
	}

	return net[0], nil
}

// List returns a list of NetworkTypes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a NetworkType is ignored. If none of the IDs
// match a NetworkType, an error is returned.
func (cs *NetworkTypeService) List(ids []int, opts ...Option) ([]*NetworkType, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var net []*NetworkType

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &net, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get NetworkTypes with IDs %v", ids)
	}

	return net, nil
}

// Index returns an index of NetworkTypes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no NetworkTypes can
// be found using the provided options, an error is returned.
func (cs *NetworkTypeService) Index(opts ...Option) ([]*NetworkType, error) {
	var net []*NetworkType

	err := cs.client.post(cs.end, &net, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of NetworkTypes")
	}

	return net, nil
}

// Search returns a list of NetworkTypes found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no NetworkTypes are found using the provided query, an error is returned.
func (cs *NetworkTypeService) Search(qry string, opts ...Option) ([]*NetworkType, error) {
	var net []*NetworkType

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &net, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get NetworkType with query %s", qry)
	}

	return net, nil
}

// Count returns the number of NetworkTypes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which NetworkTypes to count.
func (cs *NetworkTypeService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count NetworkTypes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB NetworkType object.
func (cs *NetworkTypeService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get NetworkType fields")
	}

	return f, nil
}
