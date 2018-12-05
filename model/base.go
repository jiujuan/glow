package model

type BaseModel struct {
    ID         uint `gorm:"AUTO_INCREMENT" json:"id"`
    CreateTime int `json:"create_time"`
    UpdateTime int `json:"update_time"`
}
