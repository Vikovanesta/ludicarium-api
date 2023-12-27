package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct GameLocalization -add-tags json -w

// GameLocalization represents a language translation of a video game.
// For more information visit: https://api-docs.igdb.com/#game-localization
type GameLocalization struct {
	ID        int    `json:"id"`
	Cover     int    `json:"cover"`
	CreatedAt int    `json:"created_at"`
	Game      int    `json:"game"`
	Name      string `json:"name"`
	Region    int    `json:"region"`
	UpdatedAt int    `json:"updated_at"`
}

// GameLocalizationService handles all the API calls for the IGDB GameLocalization endpoint.
type GameLocalizationService service

// Get returns a single GameLocalization identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameLocalizations, an error is returned.
func (cs *GameLocalizationService) Get(id int, opts ...Option) (*GameLocalization, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var local []*GameLocalization

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &local, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameLocalization with ID %v", id)
	}

	return local[0], nil
}

// List returns a list of GameLocalizations identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a GameLocalization is ignored. If none of the IDs
// match a GameLocalization, an error is returned.
func (cs *GameLocalizationService) List(ids []int, opts ...Option) ([]*GameLocalization, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var local []*GameLocalization

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &local, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameLocalizations with IDs %v", ids)
	}

	return local, nil
}

// Index returns an index of GameLocalizations based solely on the provided functional
// options used to sort, filter, and paginate the results. If no GameLocalizations can
// be found using the provided options, an error is returned.
func (cs *GameLocalizationService) Index(opts ...Option) ([]*GameLocalization, error) {
	var local []*GameLocalization

	err := cs.client.post(cs.end, &local, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of GameLocalizations")
	}

	return local, nil
}

// Search returns a list of GameLocalizations found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no GameLocalizations are found using the provided query, an error is returned.
func (cs *GameLocalizationService) Search(qry string, opts ...Option) ([]*GameLocalization, error) {
	var local []*GameLocalization

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &local, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameLocalization with query %s", qry)
	}

	return local, nil
}

// Count returns the number of GameLocalizations available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which GameLocalizations to count.
func (cs *GameLocalizationService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count GameLocalizations")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB GameLocalization object.
func (cs *GameLocalizationService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get GameLocalization fields")
	}

	return f, nil
}
