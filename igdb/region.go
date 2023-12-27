package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct Region -add-tags json -w

// Region for game localization.
// For more information visit: https://api-docs.igdb.com/#region
type Region struct {
	ID         int    `json:"id"`
	Category   string `json:"category"`
	CreatedAt  int    `json:"created_at"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	UpdatedAt  int    `json:"updated_at"`
}

// RegionService handles all the API calls for the IGDB Region endpoint.
type RegionService service

// Get returns a single Region identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Regions, an error is returned.
func (cs *RegionService) Get(id int, opts ...Option) (*Region, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var reg []*Region

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.post(cs.end, &reg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Region with ID %v", id)
	}

	return reg[0], nil
}

// List returns a list of Regions identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Region is ignored. If none of the IDs
// match a Region, an error is returned.
func (cs *RegionService) List(ids []int, opts ...Option) ([]*Region, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var reg []*Region

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.post(cs.end, &reg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Regions with IDs %v", ids)
	}

	return reg, nil
}

// Index returns an index of Regions based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Regions can
// be found using the provided options, an error is returned.
func (cs *RegionService) Index(opts ...Option) ([]*Region, error) {
	var reg []*Region

	err := cs.client.post(cs.end, &reg, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Regions")
	}

	return reg, nil
}

// Count returns the number of Regions available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Regions to count.
func (cs *RegionService) Search(qry string, opts ...Option) ([]*Region, error) {
	var reg []*Region

	opts = append(opts, setSearch(qry))
	err := cs.client.post(cs.end, &reg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Regions matching \"%s\"", qry)
	}

	return reg, nil
}

// Count returns the number of Regions available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Regions to count.
func (cs *RegionService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Regions")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Region object.
func (cs *RegionService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Region fields")
	}

	return f, nil
}
