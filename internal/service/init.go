package service

import (
	"gorm.io/gorm"
)

var (
	UserService     *userService
	RoleService     *roleService
	GroupService    *groupService
	ResourceService *resourceService
	OfficeService   *officeService
)

func InitService(db *gorm.DB) {

	OfficeService = NewOfficeService(db)
	UserService = NewUserService(db)
	RoleService = NewRoleService(db)
	GroupService = NewGroupService(db)
	RoleService = NewRoleService(db)
}
