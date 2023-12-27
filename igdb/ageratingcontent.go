package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct AgeRatingContent -add-tags json -w

// AgeRatingContent is the organization behind a specific rating.
type AgeRatingContent struct {
	ID          int                      `json:"id"`
	Category    AgeRatingContentCategory `json:"category"`
	Description string                   `json:"description"`
}

// AgeRatingContentCategory specifies a regulatory organization.
type AgeRatingContentCategory int

//go:generate stringer -type=AgeRatingContentCategory

// Expected AgeRatingContentCategory enums from the IGDB.
const (
	AgeRatingContentESRB_AlcoholReference AgeRatingContentCategory = iota + 1
	AgeRatingContentESRB_AnimetedBlood
	AgeRatingContentESRB_Blood
	AgeRatingContentESRB_BloodAndGore
	AgeRatingContentESRB_CartoonViolence
	AgeRatingContentESRB_ComicMischief
	AgeRatingContentESRB_CrudeHumor
	AgeRatingContentESRB_DrugReference
	AgeRatingContentESRB_FantasyViolence
	AgeRatingContentESRB_IntenseViolence
	AgeRatingContentESRB_Language
	AgeRatingContentESRB_Lyrics
	AgeRatingContentESRB_MatureHumor
	AgeRatingContentESRB_Nudity
	AgeRatingContentESRB_PartialNudity
	AgeRatingContentESRB_RealGambling
	AgeRatingContentESRB_SexualContent
	AgeRatingContentESRB_SexualThemes
	AgeRatingContentESRB_SexualViolence
	AgeRatingContentESRB_SimulatedGambling
	AgeRatingContentESRB_StrongLanguage
	AgeRatingContentESRB_StrongLyrics
	AgeRatingContentESRB_StrongSexualContent
	AgeRatingContentESRB_SuggestiveThemes
	AgeRatingContentESRB_TobaccoReference
	AgeRatingContentESRB_UseOfAlcohol
	AgeRatingContentESRB_UseOfDrugs
	AgeRatingContentESRB_UseOfTobacco
	AgeRatingContentESRB_Violence
	AgeRatingContentESRB_ViolenceReference
	AgeRatingContentESRB_AnimatedViolence
	AgeRatingContentESRB_MildLanguage
	AgeRatingContentESRB_MildViolence
	AgeRatingContentESRB_UseOfDrugsAndAlcohol
	AgeRatingContentESRB_DrugAndAlcoholReference
	AgeRatingContentESRB_MildSuggestiveTheme
	AgeRatingContentESRB_MildCartoonViolence
	AgeRatingContentESRB_MildBlood
	AgeRatingContentESRB_RealisticBloodAndGore
	AgeRatingContentESRB_RealisticViolence
	AgeRatingContentESRB_AlcoholAndTobaccoReference
	AgeRatingContentESRB_MatureSexualTheme
	AgeRatingContentESRB_MildAnimatedViolence
	AgeRatingContentESRB_MildSexualTheme
	AgeRatingContentESRB_UseOfAlcoholAndTobacco
	AgeRatingContentESRB_AnimatedBloodAndGore
	AgeRatingContentESRB_MildFantasyViolence
	AgeRatingContentESRB_MildLyrics
	AgeRatingContentESRB_RealisticBlood
	AgeRatingContentPEGI_Violence
	AgeRatingContentPEGI_Discrimination
	AgeRatingContentPEGI_BadLanguage
	AgeRatingContentPEGI_Gambling
	AgeRatingContentPEGI_OnlineGameplay
	AgeRatingContentPEGI_InGamePurchases
	AgeRatingContentCERO_Love
	AgeRatingContentCERO_SexualContent
	AgeRatingContentCERO_Violence
	AgeRatingContentCERO_Horror
	AgeRatingContentCERO_DrinkingSmoking
	AgeRatingContentCERO_Gambling
	AgeRatingContentCERO_Crime
	AgeRatingContentCERO_ControlledSubstances
	AgeRatingContentCERO_LanguageAndOthers
	AgeRatingContentGRAC_Sexuality
	AgeRatingContentGRAC_Violence
	AgeRatingContentGRAC_FearHorrorThreatening
	AgeRatingContentGRAC_Languange
	AgeRatingContentGRAC_AlcoholTobaccoDrug
	AgeRatingContentGRAC_CrimeAntiSocial
	AgeRatingContentGRAC_Gambling
	AgeRatingContentCLASS_IND_Violencia
	AgeRatingContentCLASS_IND_ViolenciaExtrema
	AgeRatingContentCLASS_IND_ConteudoSexual
	AgeRatingContentCLASS_IND_Nudez
	AgeRatingContentCLASS_IND_Sexo
	AgeRatingContentCLASS_IND_SexoExplicito
	AgeRatingContentCLASS_IND_Drogas
	AgeRatingContentCLASS_IND_DrogasLicitas
	AgeRatingContentCLASS_IND_DrogasIlícitas
	AgeRatingContentCLASS_IND_LinguagemImpropria
	AgeRatingContentCLASS_IND_AtosCriminosos
)

// AgeRatingContentService handles all the API calls for the IGDB AgeRatingContent endpoint.
type AgeRatingContentService service

// Get returns a single AgeRatingContent identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any AgeRatingContents, an error is returned.
func (as *AgeRatingContentService) Get(id int, opts ...Option) (*AgeRatingContent, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var cont []*AgeRatingContent

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := as.client.post(as.end, &cont, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AgeRatingContent with ID %v", id)
	}

	return cont[0], nil
}

// List returns a list of AgeRatingContents identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a AgeRatingContent is ignored. If none of the IDs
// match a AgeRatingContent, an error is returned.
func (as *AgeRatingContentService) List(ids []int, opts ...Option) ([]*AgeRatingContent, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var cont []*AgeRatingContent

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := as.client.post(as.end, &cont, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AgeRatingContents with IDs %v", ids)
	}

	return cont, nil
}

// Index returns an index of AgeRatingContents based solely on the provided functional
// options used to sort, filter, and paginate the results. If no AgeRatingContents can
// be found using the provided options, an error is returned.
func (as *AgeRatingContentService) Index(opts ...Option) ([]*AgeRatingContent, error) {
	var cont []*AgeRatingContent

	err := as.client.post(as.end, &cont, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of AgeRatingContents")
	}

	return cont, nil
}

// Count returns the number of AgeRatingContents available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which AgeRatingContents to count.
func (as *AgeRatingContentService) Count(opts ...Option) (int, error) {
	ct, err := as.client.getCount(as.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count AgeRatingContents")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB AgeRatingContent object.
func (as *AgeRatingContentService) Fields() ([]string, error) {
	f, err := as.client.getFields(as.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get AgeRatingContent fields")
	}

	return f, nil
}
