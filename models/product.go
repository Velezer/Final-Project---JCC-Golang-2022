package models

type Product struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Count uint   `json:"count"`

	StoreId uint  `json:"store_id"`
	Store   Store `json:"-"`

	Categories []Category `json:"-" gorm:"many2many:product_category;"`
}
