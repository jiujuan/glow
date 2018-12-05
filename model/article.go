package model

type Articles struct {
    BaseModel
    Title       string `json:"title"`
    CID         int    `json:"cid"`
    Summary     string `json:"summary"`
    Content     string `json:"content"`
    HtmlContent string `json:"html_content"`
    Author      string `json:"author"`
    Username    string `json:"username"`
    ViewCount   int    `json:"view_count"`
    Status      uint   `json:"status"`
}

const (
    ArticleIsPublished = 1  //文章已发布
    ArticleIsPublish   = 0  //未发布
    ArticleTempStore   = 2  //文章暂存
)
