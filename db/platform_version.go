package db

type PlatformVersion struct {
	Model
	Name                        string                       `gorm:"not null" json:"name"`
	Slug                        string                       `gorm:"not null" json:"slug"`
	Summary                     string                       `json:"summary"`
	MainManufacturerID          uint                         `json:"-"`
	MainManufacturer            PlatformVersionCompany       `gorm:"foreignKey:MainManufacturerID" json:"main_manufacturer"`
	CPU                         string                       `json:"cpu"`
	Graphics                    string                       `json:"graphics"`
	Storage                     string                       `json:"storage"`
	Memory                      string                       `json:"memory"`
	OS                          string                       `json:"os"`
	Resolutions                 string                       `json:"resolutions"`
	Connectivity                string                       `json:"connectivity"`
	Media                       string                       `json:"media"`
	Output                      string                       `json:"output"`
	Sound                       string                       `json:"sound"`
	LogoID                      uint                         `json:"-"`
	Logo                        PlatformLogo                 `gorm:"foreignKey:LogoID" json:"logo"`
	PlatformVersionReleaseDates []PlatformVersionReleaseDate `gorm:"foreignKey:PlatformVersionID" json:"platform_version_release_dates"`
	Companies                   []PlatformVersionCompany     `gorm:"many2many:platform_platform_version_company;" json:"companies"`
}
