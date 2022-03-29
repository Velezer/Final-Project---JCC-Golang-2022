package models

type User struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	RoleId   uint   `json:"role_id"`
	Role     Role   `json:"-"`
}
