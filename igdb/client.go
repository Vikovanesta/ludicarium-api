package igdb

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Henry-Sarabia/apicalypse"
	"github.com/pkg/errors"
)

const igdbURL string = "https://api.igdb.com/v4/"

type service struct {
	client *Client
	end    endpoint
}

type Client struct {
	http     *http.Client
	rootURL  string
	clientID string
	token    string

	AgeRatings                  *AgeRatingService
	AgeRatingContents           *AgeRatingContentService
	AlternativeNames            *AlternativeNameService
	Artworks                    *ArtworkService
	Characters                  *CharacterService
	CharacterMugshots           *CharacterMugshotService
	Collections                 *CollectionService
	CollectionMemberships       *CollectionMembershipService
	CollectionMembershipTypes   *CollectionMembershipTypeService
	CollectionRelations         *CollectionRelationService
	CollectionRelationTypes     *CollectionRelationTypeService
	CollectionTypes             *CollectionTypeService
	Companies                   *CompanyService
	CompanyLogos                *CompanyLogoService
	CompanyWebsites             *CompanyWebsiteService
	Covers                      *CoverService
	Events                      *EventService
	EventLogos                  *EventLogoService
	EventNetworks               *EventNetworkService
	ExternalGames               *ExternalGameService
	Franchises                  *FranchiseService
	Games                       *GameService
	GameEngines                 *GameEngineService
	GameEngineLogos             *GameEngineLogoService
	GameLocalizations           *GameLocalizationService
	GameModes                   *GameModeService
	GameVersions                *GameVersionService
	GameVersionFeatures         *GameVersionFeatureService
	GameVersionFeatureValues    *GameVersionFeatureValueService
	GameVideos                  *GameVideoService
	Genres                      *GenreService
	InvolvedCompanies           *InvolvedCompanyService
	Keywords                    *KeywordService
	Languages                   *LanguageService
	LanguageSupports            *LanguageSupportService
	LanguageSupportTypes        *LanguageSupportTypeService
	MultiplayerModes            *MultiplayerModeService
	NetworkTypes                *NetworkTypeService
	Platforms                   *PlatformService
	PlatformFamilies            *PlatformFamilyService
	PlatformLogos               *PlatformLogoService
	PlatformVersions            *PlatformVersionService
	PlatformVersionCompanies    *PlatformVersionCompanyService
	PlatformVersionReleaseDates *PlatformVersionReleaseDateService
	PlatformWebsites            *PlatformWebsiteService
	PlayerPerspectives          *PlayerPerspectiveService
	Regions                     *RegionService
	ReleaseDates                *ReleaseDateService
	ReleaseDateStatuses         *ReleaseDateStatusService
	Screenshots                 *ScreenshotService
	Themes                      *ThemeService
	Websites                    *WebsiteService
}

func NewClient(clientID string, appAccessToken string, custom *http.Client) *Client {
	if custom == nil {
		custom = http.DefaultClient
	}

	c := &Client{
		http:     custom,
		rootURL:  igdbURL,
		clientID: clientID,
		token:    appAccessToken,
	}

	c.AgeRatings = &AgeRatingService{client: c, end: EndpointAgeRating}
	c.AgeRatingContents = &AgeRatingContentService{client: c, end: EndpointAgeRatingContent}
	c.AlternativeNames = &AlternativeNameService{client: c, end: EndpointAlternativeName}
	c.Artworks = &ArtworkService{client: c, end: EndpointArtwork}
	c.Characters = &CharacterService{client: c, end: EndpointCharacter}
	c.CharacterMugshots = &CharacterMugshotService{client: c, end: EndpointCharacterMugshot}
	c.Collections = &CollectionService{client: c, end: EndpointCollection}
	c.CollectionMemberships = &CollectionMembershipService{client: c, end: EndpointCollectionMembership}
	c.CollectionMembershipTypes = &CollectionMembershipTypeService{client: c, end: EndpointCollectionMembershipType}
	c.CollectionRelations = &CollectionRelationService{client: c, end: EndpointCollectionRelation}
	c.CollectionRelationTypes = &CollectionRelationTypeService{client: c, end: EndpointCollectionRelationType}
	c.CollectionTypes = &CollectionTypeService{client: c, end: EndpointCollectionType}
	c.Companies = &CompanyService{client: c, end: EndpointCompany}
	c.CompanyLogos = &CompanyLogoService{client: c, end: EndpointCompanyLogo}
	c.CompanyWebsites = &CompanyWebsiteService{client: c, end: EndpointCompanyWebsite}
	c.Covers = &CoverService{client: c, end: EndpointCover}
	c.Events = &EventService{client: c, end: EndpointEvent}
	c.EventLogos = &EventLogoService{client: c, end: EndpointEventLogo}
	c.EventNetworks = &EventNetworkService{client: c, end: EndpointEventNetwork}
	c.ExternalGames = &ExternalGameService{client: c, end: EndpointExternalGame}
	c.Franchises = &FranchiseService{client: c, end: EndpointFranchise}
	c.Games = &GameService{client: c, end: EndpointGame}
	c.GameEngines = &GameEngineService{client: c, end: EndpointGameEngine}
	c.GameEngineLogos = &GameEngineLogoService{client: c, end: EndpointGameEngineLogo}
	c.GameLocalizations = &GameLocalizationService{client: c, end: EndpointGameLocalization}
	c.GameModes = &GameModeService{client: c, end: EndpointGameMode}
	c.GameVersions = &GameVersionService{client: c, end: EndpointGameVersion}
	c.GameVersionFeatures = &GameVersionFeatureService{client: c, end: EndpointGameVersionFeature}
	c.GameVersionFeatureValues = &GameVersionFeatureValueService{client: c, end: EndpointGameVersionFeatureValue}
	c.GameVideos = &GameVideoService{client: c, end: EndpointGameVideo}
	c.Genres = &GenreService{client: c, end: EndpointGenre}
	c.InvolvedCompanies = &InvolvedCompanyService{client: c, end: EndpointInvolvedCompany}
	c.Keywords = &KeywordService{client: c, end: EndpointKeyword}
	c.Languages = &LanguageService{client: c, end: EndpointLanguage}
	c.LanguageSupports = &LanguageSupportService{client: c, end: EndpointLanguageSupport}
	c.LanguageSupportTypes = &LanguageSupportTypeService{client: c, end: EndpointLanguageSupportType}
	c.MultiplayerModes = &MultiplayerModeService{client: c, end: EndpointMultiplayerMode}
	c.NetworkTypes = &NetworkTypeService{client: c, end: EndpointNetworkType}
	c.Platforms = &PlatformService{client: c, end: EndpointPlatform}
	c.PlatformFamilies = &PlatformFamilyService{client: c, end: EndpointPlatformFamily}
	c.PlatformLogos = &PlatformLogoService{client: c, end: EndpointPlatformLogo}
	c.PlatformVersions = &PlatformVersionService{client: c, end: EndpointPlatformVersion}
	c.PlatformVersionCompanies = &PlatformVersionCompanyService{client: c, end: EndpointPlatformVersionCompany}
	c.PlatformVersionReleaseDates = &PlatformVersionReleaseDateService{client: c, end: EndpointPlatformVersionReleaseDate}
	c.PlatformWebsites = &PlatformWebsiteService{client: c, end: EndpointPlatformWebsite}
	c.PlayerPerspectives = &PlayerPerspectiveService{client: c, end: EndpointPlayerPerspective}
	c.Regions = &RegionService{client: c, end: EndpointRegion}
	c.ReleaseDates = &ReleaseDateService{client: c, end: EndpointReleaseDate}
	c.ReleaseDateStatuses = &ReleaseDateStatusService{client: c, end: EndpointReleaseDateStatus}
	c.Screenshots = &ScreenshotService{client: c, end: EndpointScreenshot}
	c.Themes = &ThemeService{client: c, end: EndpointTheme}
	c.Websites = &WebsiteService{client: c, end: EndpointWebsite}

	return c
}

// Request configures a new request for the provided URL and
// adds the necessary headers to communicate with the IGDB.
func (c *Client) request(end endpoint, opts ...Option) (*http.Request, error) {
	unwrapped, err := unwrapOptions(opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create request with invalid options")
	}

	req, err := apicalypse.NewRequest("POST", c.rootURL+string(end), unwrapped...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot make request for '%s' endpoint", end)
	}

	req.Header.Add("client-id", c.clientID)
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("x-user-agent", "HenrySarabia/igdb")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// Send sends the provided request and stores the response in the value pointed to by result.
// The response will be checked and return any errors.
func (c *Client) send(req *http.Request, result interface{}) error {
	resp, err := c.http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client cannot send request")
	}
	defer resp.Body.Close()

	if err = checkResponse(resp); err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "cannot read response body")
	}

	if isBracketPair(b) {
		return ErrNoResults
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return errors.Wrap(errInvalidJSON, err.Error())
	}

	return nil
}

// post sends a POST request to the provided endpoint with the provided options and
// stores the results in the value pointed to by result.
func (c *Client) post(end endpoint, result interface{}, opts ...Option) error {
	req, err := c.request(end, opts...)
	if err != nil {
		return err
	}

	err = c.send(req, result)
	if err != nil {
		return errors.Wrap(err, "cannot make POST request")
	}

	return nil
}
