package model

import "time"
//机构信息  
type Group struct { 
    //
    Id  int64  `gorm:"column:id;PRIMARY_KEY" json:"Id,omitempty"` 
    //组代号
    Code  string  `gorm:"column:code" json:"Code,omitempty"` 
    //组名称
    Name  string  `gorm:"column:name" json:"Name,omitempty"` 
    //
    ParentId  int64  `gorm:"column:parent_id" json:"ParentId,omitempty"` 
    //
    CreatedAt  time.Time  `gorm:"column:created_at" json:"CreatedAt,omitempty"` 
    //
    UpdatedAt  time.Time  `gorm:"column:updated_at" json:"UpdatedAt,omitempty"`
}

func (u Group) TableName() string {
	return "public.sys_group"
}
