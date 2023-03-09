package model

import "time"
//角色信息  
type Role struct { 
    //
    Id  int64  `gorm:"column:id;PRIMARY_KEY" json:"Id,omitempty"` 
    //角色代号
    Code  string  `gorm:"column:code" json:"Code,omitempty"` 
    //角色名称
    Name  string  `gorm:"column:name" json:"Name,omitempty"` 
    //
    GroupId  int64  `gorm:"column:group_id" json:"GroupId,omitempty"` 
    //
    CreatedAt  time.Time  `gorm:"column:created_at" json:"CreatedAt,omitempty"` 
    //
    UpdatedAt  time.Time  `gorm:"column:updated_at" json:"UpdatedAt,omitempty"`
}

func (u Role) TableName() string {
	return "public.sys_role"
}
