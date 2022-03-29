package models

type Store struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`

	UserId uint `json:"user_id"`
	User   User `json:"-"`

	Products []Product `json:"-"`
}
