package model

type Category struct {
    BaseModel
    CatName string `json:"cat_name"`
    CatInfo string `json:"cat_info"`
    CatKey  string `json:"cat_key"`
}
