package controllers

import "hewantani/services"

type Controller struct {
	UserService  services.UserIface
	RoleService  services.RoleIface
	StoreService services.StoreIface
}
