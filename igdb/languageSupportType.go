package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct LanguageSupportType -add-tags json -w

// LanguageSupportType contains the identifiers for the support types that language support uses.
// For more information visit: https://api-docs.igdb.com/#language-support-type
type LanguageSupportType struct {
	ID        int    `json:"id"`
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	UpdatedAt int    `json:"updated_at"`
}

// LanguageSupportTypeService handles all the API calls for the IGDB LanguageSupportType endpoint.
type LanguageSupportTypeService service

// Get returns a single LanguageSupportType identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any LanguageSupportTypes, an error is returned.
func (cs *LanguageSupportTypeService) Get(id int, opts ...Option) (*LanguageSupportType, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var sup []*LanguageSupportType

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &sup, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get LanguageSupportType with ID %v", id)
	}

	return sup[0], nil
}

// List returns a list of LanguageSupportTypes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a LanguageSupportType is ignored. If none of the IDs
// match a LanguageSupportType, an error is returned.
func (cs *LanguageSupportTypeService) List(ids []int, opts ...Option) ([]*LanguageSupportType, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var sup []*LanguageSupportType

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &sup, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get LanguageSupportTypes with IDs %v", ids)
	}

	return sup, nil
}

// Index returns an index of LanguageSupportTypes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no LanguageSupportTypes can
// be found using the provided options, an error is returned.
func (cs *LanguageSupportTypeService) Index(opts ...Option) ([]*LanguageSupportType, error) {
	var sup []*LanguageSupportType

	err := cs.client.post(cs.end, &sup, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of LanguageSupportTypes")
	}

	return sup, nil
}

// Search returns a list of LanguageSupportTypes found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no LanguageSupportTypes are found using the provided query, an error is returned.
func (cs *LanguageSupportTypeService) Search(qry string, opts ...Option) ([]*LanguageSupportType, error) {
	var sup []*LanguageSupportType

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &sup, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get LanguageSupportType with query %s", qry)
	}

	return sup, nil
}

// Count returns the number of LanguageSupportTypes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which LanguageSupportTypes to count.
func (cs *LanguageSupportTypeService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count LanguageSupportTypes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB LanguageSupportType object.
func (cs *LanguageSupportTypeService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get LanguageSupportType fields")
	}

	return f, nil
}
