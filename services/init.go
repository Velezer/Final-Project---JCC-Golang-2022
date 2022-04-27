package services

import "gorm.io/gorm"

type all struct {
	UserService        UserIface
	RoleService        RoleIface
	OrderStatusService OrderStatusIface
	StoreService       StoreIface
	ProductService     ProductIface
	CategoryService    CategoryIface
	CartService        CartIface
	OrderService       OrderIface
}

var All *all

func Setup(Db *gorm.DB) {
	if All == nil {
		All = &all{}
		All.UserService = &User{Db}
		All.RoleService = &Role{Db}
		All.StoreService = &Store{Db}
		All.ProductService = &Product{Db}
		All.CategoryService = &Category{Db}
		All.CartService = &Cart{Db}
		All.OrderService = &Order{Db}
		All.OrderStatusService = &OrderStatus{Db}
	}

}
