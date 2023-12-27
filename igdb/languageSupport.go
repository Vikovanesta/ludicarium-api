package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct LanguageSupport -add-tags json -w

// LanguageSupport represents a language support for a video game.
// For more information visit: https://api-docs.igdb.com/#language-support
type LanguageSupport struct {
	ID                  int `json:"id"`
	CreatedAt           int `json:"created_at"`
	Game                int `json:"game"`
	Language            int `json:"language"`
	LanguageSupportType int `json:"language_support_type"`
	UpdatedAt           int `json:"updated_at"`
}

// LanguageSupportService handles all the API calls for the IGDB LanguageSupport endpoint.
type LanguageSupportService service

// Get returns a single LanguageSupport identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any LanguageSupports, an error is returned.
func (cs *LanguageSupportService) Get(id int, opts ...Option) (*LanguageSupport, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var sup []*LanguageSupport

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &sup, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get LanguageSupport with ID %v", id)
	}

	return sup[0], nil
}

// List returns a list of LanguageSupports identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a LanguageSupport is ignored. If none of the IDs
// match a LanguageSupport, an error is returned.
func (cs *LanguageSupportService) List(ids []int, opts ...Option) ([]*LanguageSupport, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var sup []*LanguageSupport

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &sup, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get LanguageSupports with IDs %v", ids)
	}

	return sup, nil
}

// Index returns an index of LanguageSupports based solely on the provided functional
// options used to sort, filter, and paginate the results. If no LanguageSupports can
// be found using the provided options, an error is returned.
func (cs *LanguageSupportService) Index(opts ...Option) ([]*LanguageSupport, error) {
	var sup []*LanguageSupport

	err := cs.client.post(cs.end, &sup, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of LanguageSupports")
	}

	return sup, nil
}

// Search returns a list of LanguageSupports found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no LanguageSupports are found using the provided query, an error is returned.
func (cs *LanguageSupportService) Search(qry string, opts ...Option) ([]*LanguageSupport, error) {
	var sup []*LanguageSupport

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &sup, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get LanguageSupport with query %s", qry)
	}

	return sup, nil
}

// Count returns the number of LanguageSupports available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which LanguageSupports to count.
func (cs *LanguageSupportService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count LanguageSupports")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB LanguageSupport object.
func (cs *LanguageSupportService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get LanguageSupport fields")
	}

	return f, nil
}
