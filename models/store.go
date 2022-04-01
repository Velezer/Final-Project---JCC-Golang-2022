package models

type Store struct {
	BaseModel

	Name        string `json:"name" gorm:"not null;unique;check:name <> ''"`
	Description string `json:"description"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`

	UserId uint `json:"user_id"`
	User   User `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Products []Product `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
