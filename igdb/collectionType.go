package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct CollectionType -add-tags json -w

// CollectionType represents a video game series.
// For more information visit: https://api-docs.igdb.com/#collection-type
type CollectionType struct {
	ID          int    `json:"id"`
	CreatedAt   int    `json:"created_at"`
	Description string `json:"description"`
	Name        string `json:"name"`
	UpdatedAt   int    `json:"updated_at"`
}

// CollectionTypeService handles all the API calls for the IGDB CollectionType endpoint.
type CollectionTypeService service

// Get returns a single CollectionType identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any CollectionTypes, an error is returned.
func (cs *CollectionTypeService) Get(id int, opts ...Option) (*CollectionType, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ct []*CollectionType

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &ct, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionType with ID %v", id)
	}

	return ct[0], nil
}

// List returns a list of CollectionTypes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a CollectionType is ignored. If none of the IDs
// match a CollectionType, an error is returned.
func (cs *CollectionTypeService) List(ids []int, opts ...Option) ([]*CollectionType, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ct []*CollectionType

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &ct, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionTypes with IDs %v", ids)
	}

	return ct, nil
}

// Index returns an index of CollectionTypes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no CollectionTypes can
// be found using the provided options, an error is returned.
func (cs *CollectionTypeService) Index(opts ...Option) ([]*CollectionType, error) {
	var ct []*CollectionType

	err := cs.client.post(cs.end, &ct, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of CollectionTypes")
	}

	return ct, nil
}

// Search returns a list of CollectionTypes found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no CollectionTypes are found using the provided query, an error is returned.
func (cs *CollectionTypeService) Search(qry string, opts ...Option) ([]*CollectionType, error) {
	var ct []*CollectionType

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &ct, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionType with query %s", qry)
	}

	return ct, nil
}

// Count returns the number of CollectionTypes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CollectionTypes to count.
func (cs *CollectionTypeService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count CollectionTypes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB CollectionType object.
func (cs *CollectionTypeService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get CollectionType fields")
	}

	return f, nil
}
