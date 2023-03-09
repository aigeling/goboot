package rbac

import (
	"time"

	"gorm.io/gorm"
)

// 资源
type Resource struct {
	ID        uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	Code      string `gorm:"size:60;not null;comment:资源代码"`
	Name      string `gorm:"size:256;not null;comment:资源名称"`
	Obj       string `gorm:"uniqueIndex:idx_resource;size:256;not null;comment:资源对象，一般指uri"`
	Act       string `gorm:"uniqueIndex:idx_resource;size:20;not null;comment:执行操作，一般指http请求method"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Resource) TableName() string {
	return "sys_resource"
}

type Group struct {
	ID        uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	Code      string `gorm:"size:20;not null;comment:组代号"`
	Name      string `gorm:"size:256;not null;comment:组名称"`
	ParentId  uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Group) TableName() string {
	return "sys_group"
}

type Role struct {
	ID        uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	Code      string `gorm:"size:20;not null;comment:角色代号"`
	Name      string `gorm:"size:60;not null;comment:角色名称"`
	GroupID   uint   `gorm:"comment:所属机构id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Role) TableName() string {
	return "sys_role"
}

type User struct {
	ID        uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	Code      string `gorm:"size:20;"`
	LoginName string `gorm:"size:60;not null;comment:用户登录名"`
	Password  string `gorm:"size:200;comment:密码"`
	Name      string `gorm:"size:60;not null;comment:用户姓名"`
	GroupID   uint   `gorm:"comment:所属机构id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "sys_user"
}

type User_Role struct {
	UserID    uint `gorm:"primaryKey;not null"`
	RoleID    uint `gorm:"primaryKey;not null"`
	CreatedAt time.Time
}

func (User_Role) TableName() string {
	return "sys_user_role"
}

type Role_Resource struct {
	RoleID     uint `gorm:"primaryKey;not null"`
	ResourceID uint `gorm:"primaryKey;not null"`
	CreatedAt  time.Time
}

func (Role_Resource) TableName() string {
	return "sys_role_resource"
}

type Role_Group struct {
	RoleID    uint `gorm:"primaryKey;not null"`
	GroupID   uint `gorm:"primaryKey;not null"`
	CreatedAt time.Time
}

func (Role_Group) TableName() string {
	return "sys_role_group"
}

// 创建rbac模型
func CreateModel(db *gorm.DB) error {

	return db.AutoMigrate(&User{}, &Role{}, &Resource{}, &Group{}, &User_Role{}, &Role_Resource{}, &Role_Group{})
}
