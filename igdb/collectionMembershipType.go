package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

// go:generate gomodifytags -file $GOFILE -struct CollectionMembershipType -add-tags json -w

// CollectionMembershipType represents enums for collection membership types.
// For more information visit: https://api-docs.igdb.com/#collection-membership-type
type CollectionMembershipType struct {
	ID                    int    `json:"id"`
	AllowedCollectionType int    `json:"allowed_collection_type"`
	CreatedAt             int    `json:"created_at"`
	Description           string `json:"description"`
	Name                  string `json:"name"`
	UpdatedAt             int    `json:"updated_at"`
}

// CollectionMembershipTypeService handles all the API calls for the IGDB CollectionMembershipType endpoint.
type CollectionMembershipTypeService service

// Get returns a single CollectionMembershipType identified by the provided IGDB ID. Provide
// the setFields functional option if you need a subset of fields returned.
// If the ID does not match any CollectionMembershipTypes, an error is returned.
func (cs *CollectionMembershipTypeService) Get(id int, opts ...Option) (*CollectionMembershipType, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var cmt []*CollectionMembershipType

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &cmt, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionMembershipType with ID %v", id)
	}

	return cmt[0], nil
}

// List returns a list of CollectionMembershipTypes identified by the provided
// list of IGDB IDs. Provide functional options to sort, filter, and paginate the results.
// If none of the IDs match a CollectionMembershipType, an error is returned.
func (cs *CollectionMembershipTypeService) List(ids []int, opts ...Option) ([]*CollectionMembershipType, error) {
	if len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var cmt []*CollectionMembershipType

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &cmt, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionMembershipTypes with IDs %v", ids)
	}

	return cmt, nil
}

// Index returns an index of CollectionMembershipTypes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no CollectionMembershipTypes can
// be found using the provided options, an error is returned.
func (cs *CollectionMembershipTypeService) Index(opts ...Option) ([]*CollectionMembershipType, error) {
	var cmt []*CollectionMembershipType

	err := cs.client.post(cs.end, &cmt, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of CollectionMembershipTypes")
	}

	return cmt, nil
}

// Count returns the number of CollectionMembershipTypes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CollectionMembershipTypes to count.
func (cs *CollectionMembershipTypeService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count CollectionMembershipTypes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB CollectionMembershipType object.
func (cs *CollectionMembershipTypeService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get CollectionMembershipType fields")
	}

	return f, nil
}
