package controllers

import "hewantani/services"

type Controller struct {
	UserService     services.UserIface
	RoleService     services.RoleIface
	StoreService    services.StoreIface
	ProductService  services.ProductIface
	CategoryService services.CategoryIface
	CartService     services.CartIface
}
