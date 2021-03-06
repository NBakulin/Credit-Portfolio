package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Login        string
	Password     string
	Role         Role
	RoleID       uint `json:"role_id"`
	PaymentPlans []PaymentPlan
}
