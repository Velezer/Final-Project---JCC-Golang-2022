package services

import (
	"hewantani/config"

	"gorm.io/gorm"
)

type all struct {
	UserService     UserIface
	RoleService     RoleIface
	StoreService    StoreIface
	ProductService  ProductIface
	CategoryService CategoryIface
	CartService     CartIface
	OrderService    OrderIface
}

var Db *gorm.DB
var All *all

func init() {
	if Db == nil {
		Db = config.ConnectDatabase()
	}

	if All == nil {
		All = &all{}
		All.UserService = User{Db}
		All.RoleService = Role{Db}
		All.StoreService = Store{Db}
		All.ProductService = Product{Db}
		All.CategoryService = Category{Db}
		All.CartService = Cart{Db}
	}

}
