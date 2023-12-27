package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct EventLogo -add-tags json -w

// EventLogo represents logo for a gaming event.
// For more information visit: https://api-docs.igdb.com/#event-logo
type EventLogo struct {
	Image     `json:"image"`
	ID        int `json:"id"`
	CreatedAt int `json:"created_at"`
	Event     int `json:"event"`
	UpdatedAt int `json:"updated_at"`
}

// EventLogoService handles all the API calls for the IGDB EventLogo endpoint.
type EventLogoService service

// Get returns a single EventLogo identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any EventLogos, an error is returned.
func (cs *EventLogoService) Get(id int, opts ...Option) (*EventLogo, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var logo []*EventLogo

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get EventLogo with ID %v", id)
	}

	return logo[0], nil
}

// List returns a list of EventLogos identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a EventLogo is ignored. If none of the IDs
// match a EventLogo, an error is returned.
func (cs *EventLogoService) List(ids []int, opts ...Option) ([]*EventLogo, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var logo []*EventLogo

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get EventLogos with IDs %v", ids)
	}

	return logo, nil
}

// Index returns an index of EventLogos based solely on the provided functional
// options used to sort, filter, and paginate the results. If no EventLogos can
// be found using the provided options, an error is returned.
func (cs *EventLogoService) Index(opts ...Option) ([]*EventLogo, error) {
	var ct []*EventLogo

	err := cs.client.post(cs.end, &ct, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of EventLogos")
	}

	return ct, nil
}

// Search returns a list of EventLogos found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no EventLogos are found using the provided query, an error is returned.
func (cs *EventLogoService) Search(qry string, opts ...Option) ([]*EventLogo, error) {
	var logo []*EventLogo

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get EventLogo with query %s", qry)
	}

	return logo, nil
}

// Count returns the number of EventLogos available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which EventLogos to count.
func (cs *EventLogoService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count EventLogos")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB EventLogo object.
func (cs *EventLogoService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get EventLogo fields")
	}

	return f, nil
}
