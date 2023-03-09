package model

import "time"
//用户信息  
type User struct { 
    //
    Id  int64  `gorm:"column:id;PRIMARY_KEY" json:"Id,omitempty"` 
    //
    Code  string  `gorm:"column:code" json:"Code,omitempty"` 
    //用户登录名
    LoginName  string  `gorm:"column:login_name" json:"LoginName,omitempty"` 
    //密码
    Password  string  `gorm:"column:password" json:"Password,omitempty"` 
    //用户姓名
    Name  string  `gorm:"column:name" json:"Name,omitempty"` 
    //
    GroupId  int64  `gorm:"column:group_id" json:"GroupId,omitempty"` 
    //
    CreatedAt  time.Time  `gorm:"column:created_at" json:"CreatedAt,omitempty"` 
    //
    UpdatedAt  time.Time  `gorm:"column:updated_at" json:"UpdatedAt,omitempty"`
}

func (u User) TableName() string {
	return "public.sys_user"
}
