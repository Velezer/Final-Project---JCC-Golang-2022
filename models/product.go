package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name     string `json:"name"`
	Count    uint   `json:"count"`
	Price    uint   `json:"price"`
	ImageUrl string `json:"image_url"`

	StoreId uint  `json:"store_id"`
	Store   Store `json:"-"`

	Categories []Category `json:"-" gorm:"many2many:product_category;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
