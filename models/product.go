package models

type Product struct {
	BaseModel

	Name        string `json:"name"`
	Description string `json:"description"`
	Count       uint   `json:"count"`
	Price       uint   `json:"price"`
	ImageUrl    string `json:"image_url"`

	StoreId uint  `json:"store_id"`
	Store   Store `json:"-"`

	UserId uint `json:"user_id"`
	User   User `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Categories []Category `json:"-" gorm:"many2many:product_category;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
