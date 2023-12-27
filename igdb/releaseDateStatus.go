package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct ReleaseDateStatus -add-tags json -w

// ReleaseDateStatus provides definition of all of the current release date statuses.
// For more information visit: https://api-docs.igdb.com/#release-date-status
type ReleaseDateStatus struct {
	ID          int    `json:"id"`
	CreatedAt   int    `json:"created_at"`
	Description string `json:"description"`
	Name        string `json:"name"`
	UpdatedAt   int    `json:"updated_at"`
}

// ReleaseDateStatusService handles all the API calls for the IGDB ReleaseDateStatus endpoint.
type ReleaseDateStatusService service

// Get returns a single ReleaseDateStatus identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any ReleaseDateStatuss, an error is returned.
func (cs *ReleaseDateStatusService) Get(id int, opts ...Option) (*ReleaseDateStatus, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var stat []*ReleaseDateStatus

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &stat, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ReleaseDateStatus with ID %v", id)
	}

	return stat[0], nil
}

// List returns a list of ReleaseDateStatuss identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a ReleaseDateStatus is ignored. If none of the IDs
// match a ReleaseDateStatus, an error is returned.
func (cs *ReleaseDateStatusService) List(ids []int, opts ...Option) ([]*ReleaseDateStatus, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var stat []*ReleaseDateStatus

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &stat, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ReleaseDateStatuss with IDs %v", ids)
	}

	return stat, nil
}

// Index returns an index of ReleaseDateStatuss based solely on the provided functional
// options used to sort, filter, and paginate the results. If no ReleaseDateStatuss can
// be found using the provided options, an error is returned.
func (cs *ReleaseDateStatusService) Index(opts ...Option) ([]*ReleaseDateStatus, error) {
	var stat []*ReleaseDateStatus

	err := cs.client.post(cs.end, &stat, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of ReleaseDateStatuss")
	}

	return stat, nil
}

// Count returns the number of ReleaseDateStatuss available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which ReleaseDateStatuss to count.
func (cs *ReleaseDateStatusService) Search(qry string, opts ...Option) ([]*ReleaseDateStatus, error) {
	var stat []*ReleaseDateStatus

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &stat, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ReleaseDateStatuss matching \"%s\"", qry)
	}

	return stat, nil
}

// Count returns the number of ReleaseDateStatuss available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which ReleaseDateStatuss to count.
func (cs *ReleaseDateStatusService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count ReleaseDateStatuss")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB ReleaseDateStatus object.
func (cs *ReleaseDateStatusService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ReleaseDateStatus fields")
	}

	return f, nil
}
