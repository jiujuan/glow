package model

type SysConfig struct {
    BaseModel
    SiteConfig string
    SiteName   string
    KeyWord    string
    SiteDesc   string
    IsOpen     int   `gorm:"default:1"`
    CloseDesc  string
    AdminEmail string
}