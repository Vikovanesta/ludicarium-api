package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct Language -add-tags json -w

// Language that are used in languange support endpoint.
// For more information visit: https://api-docs.igdb.com/#language
type Language struct {
	ID         int    `json:"id"`
	CreatedAt  int    `json:"created_at"`
	Locale     string `json:"locale"`
	Name       string `json:"name"`
	NativeName string `json:"native_name"`
	UpdatedAt  int    `json:"updated_at"`
}

// LanguageService handles all the API calls for the IGDB Language endpoint.
type LanguageService service

// Get returns a single Language identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Languages, an error is returned.
func (cs *LanguageService) Get(id int, opts ...Option) (*Language, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var lang []*Language

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &lang, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Language with ID %v", id)
	}

	return lang[0], nil
}

// List returns a list of Languages identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Language is ignored. If none of the IDs
// match a Language, an error is returned.
func (cs *LanguageService) List(ids []int, opts ...Option) ([]*Language, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var lang []*Language

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &lang, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Languages with IDs %v", ids)
	}

	return lang, nil
}

// Index returns an index of Languages based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Languages can
// be found using the provided options, an error is returned.
func (cs *LanguageService) Index(opts ...Option) ([]*Language, error) {
	var lang []*Language

	err := cs.client.post(cs.end, &lang, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Languages")
	}

	return lang, nil
}

// Search returns a list of Languages found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Languages are found using the provided query, an error is returned.
func (cs *LanguageService) Search(qry string, opts ...Option) ([]*Language, error) {
	var lang []*Language

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &lang, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Language with query %s", qry)
	}

	return lang, nil
}

// Count returns the number of Languages available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Languages to count.
func (cs *LanguageService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Languages")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Language object.
func (cs *LanguageService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Language fields")
	}

	return f, nil
}
