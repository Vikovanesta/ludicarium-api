package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct CollectionRelation -add-tags json -w

// CollectionRelation represents a relationship between collections.
// For more information visit: https://api-docs.igdb.com/#collection-relation
type CollectionRelation struct {
	ID               int `json:"id"`
	ChildCollection  int `json:"child_collection"`
	CreatedAt        int `json:"created_at"`
	ParentCollection int `json:"parent_collection"`
	Type             int `json:"type"`
	UpdatedAt        int `json:"updated_at"`
}

// CollectionRelationService handles all the API calls for the IGDB CollectionRelation endpoint.
type CollectionRelationService service

// Get returns a single CollectionRelation identified by the provided IGDB ID. Provide
// the setFields functional option if you need a subset of data returned.
// If the ID does not match any CollectionRelations, an error is returned.
func (cr *CollectionRelationService) Get(id int, opts ...Option) (*CollectionRelation, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var rel []*CollectionRelation

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cr.client.post(cr.end, &rel, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionRelation with ID %v", id)
	}

	return rel[0], nil
}

// List returns a list of CollectionRelations identified by the provided
// list of IGDB IDs. Provide the SetFields functional option if you need
// a subset of data returned. If any IDs are invalid, the whole request
// will fail and an error will be returned.
func (cr *CollectionRelationService) List(ids []int, opts ...Option) ([]*CollectionRelation, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var rel []*CollectionRelation

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cr.client.post(cr.end, &rel, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionRelations with IDs %v", ids)
	}

	return rel, nil
}

// Index returns an index of CollectionRelations based solely on the provided functional
// options used to sort, filter, and paginate the results. If no CollectionRelation can
// be found using the provided options, an error is returned.
func (cr *CollectionRelationService) Index(opts ...Option) ([]*CollectionRelation, error) {
	var rel []*CollectionRelation

	err := cr.client.post(cr.end, &rel, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of CollectionRelations")
	}

	return rel, nil
}

// Search returns a list of CollectionRelations found by searching the IGDB using the provided
// query. Provide the SetFields functional option if you need a subset of data returned.
// If no CollectionRelations are found using the provided query, an error is returned.
func (cr *CollectionRelationService) Search(qry string, opts ...Option) ([]*CollectionRelation, error) {
	var rel []*CollectionRelation

	opts = append(opts, setSearch(qry))
	err := cr.client.post(cr.end, &rel, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CollectionRelations matching \"%s\"", qry)
	}

	return rel, nil
}

// Count returns the number of CollectionRelations available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CollectionRelations to count.
func (cr *CollectionRelationService) Count(opts ...Option) (int, error) {
	ct, err := cr.client.getCount(cr.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count CollectionRelations")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB CollectionRelation object.
func (cr *CollectionRelationService) Fields() ([]string, error) {
	f, err := cr.client.getFields(cr.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get CollectionRelation fields")
	}

	return f, nil
}
