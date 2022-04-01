package models

type Category struct {
	BaseModel

	Name string `json:"name" gorm:"not null;unique"`

	Product []Product `json:"-" gorm:"many2many:product_category;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
