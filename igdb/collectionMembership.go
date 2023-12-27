package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct CollectionMembership -add-tags json -w

// CollectionMembership represents a game's membership to a collection.
// For more information visit: https://api-docs.igdb.com/#collection-membership
type CollectionMembership struct {
	ID         int `json:"id"`
	Collection int `json:"collection"`
	CreatedAt  int `json:"created_at"`
	Game       int `json:"game"`
	Type       int `json:"type"`
	UpdatedAt  int `json:"updated_at"`
}

// CollectionMembershipService handles all the API calls for the IGDB CollectionMembership endpoint.
type CollectionMembershipService service

// Get returns a single CollectionMembership identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any CollectionMemberships, an error is returned.
func (cs *CollectionMembershipService) Get(id int, opts ...Option) (*CollectionMembership, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var mem []*CollectionMembership

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &mem, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionMembership with ID %v", id)
	}

	return mem[0], nil
}

// List returns a list of CollectionMemberships identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a CollectionMembership is ignored. If none of the IDs
// match a CollectionMembership, an error is returned.
func (cs *CollectionMembershipService) List(ids []int, opts ...Option) ([]*CollectionMembership, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var mem []*CollectionMembership

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &mem, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionMemberships with IDs %v", ids)
	}

	return mem, nil
}

// Index returns an index of CollectionMemberships based solely on the provided functional
// options used to sort, filter, and paginate the results. If no CollectionMemberships can
// be found using the provided options, an error is returned.
func (cs *CollectionMembershipService) Index(opts ...Option) ([]*CollectionMembership, error) {
	var mem []*CollectionMembership

	err := cs.client.post(cs.end, &mem, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of CollectionMemberships")
	}

	return mem, nil
}

// Count returns the number of CollectionMemberships available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CollectionMemberships to count.
func (cs *CollectionMembershipService) Search(qry string, opts ...Option) ([]*CollectionMembership, error) {
	var mem []*CollectionMembership

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &mem, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionMemberships matching \"%s\"", qry)
	}

	return mem, nil
}

// Count returns the number of CollectionMemberships available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CollectionMemberships to count.
func (cs *CollectionMembershipService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count CollectionMemberships")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB CollectionMembership object.
func (cs *CollectionMembershipService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get CollectionMembership fields")
	}

	return f, nil
}
