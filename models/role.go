package models

type Role struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"not null;unique"`
}