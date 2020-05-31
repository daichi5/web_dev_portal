package models

type User struct {
	CommonModel
	Name string `gorm:"size:255" json:"name"`
}
