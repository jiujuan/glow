package model

type Comments struct {
    BaseModel
    Nickname  string `json:"nickname"`
    SiteUrl   string `json:"site_url"`
    Email     string `json:"email"`
    ArticleId uint   `json:"article_id"`
}
