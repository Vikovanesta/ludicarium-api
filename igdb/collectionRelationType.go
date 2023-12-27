package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct CollectionRelationType -add-tags json -w

// CollectionRelationType represents the type of relationship between collections.
// For more information visit: https://api-docs.igdb.com/#collection-relation-type
type CollectionRelationType struct {
	ID                int    `json:"id"`
	AllowedChildType  int    `json:"allowed_child_type"`
	AllowedParentType int    `json:"allowed_parent_type"`
	CreatedAt         int    `json:"created_at"`
	Description       string `json:"description"`
	Name              string `json:"name"`
	UpdatedAt         int    `json:"updated_at"`
}

// CollectionRelationTypeService handles all the API calls for the IGDB CollectionRelationType endpoint.
type CollectionRelationTypeService service

// Get returns a single CollectionRelationType identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any CollectionRelationTypes, an error is returned.
func (cs *CollectionRelationTypeService) Get(id int, opts ...Option) (*CollectionRelationType, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var crt []*CollectionRelationType

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &crt, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionRelationType with ID %v", id)
	}

	return crt[0], nil
}

// List returns a list of CollectionRelationTypes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a CollectionRelationType is ignored. If none of the IDs
// match a CollectionRelationType, an error is returned.
func (cs *CollectionRelationTypeService) List(ids []int, opts ...Option) ([]*CollectionRelationType, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var crt []*CollectionRelationType

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &crt, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionRelationTypes with IDs %v", ids)
	}

	return crt, nil
}

// Index returns an index of CollectionRelationTypes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no CollectionRelationTypes can
// be found using the provided options, an error is returned.
func (cs *CollectionRelationTypeService) Index(opts ...Option) ([]*CollectionRelationType, error) {
	var crt []*CollectionRelationType

	err := cs.client.post(cs.end, &crt, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of CollectionRelationTypes")
	}

	return crt, nil
}

// Search returns a list of CollectionRelationTypes found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no CollectionRelationTypes are found using the provided query, an error is returned.
func (cs *CollectionRelationTypeService) Search(qry string, opts ...Option) ([]*CollectionRelationType, error) {
	var crt []*CollectionRelationType

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &crt, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionRelationType with query %s", qry)
	}

	return crt, nil
}

// Count returns the number of CollectionRelationTypes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CollectionRelationTypes to count.
func (cs *CollectionRelationTypeService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count CollectionRelationTypes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB CollectionRelationType object.
func (cs *CollectionRelationTypeService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get CollectionRelationType fields")
	}

	return f, nil
}
