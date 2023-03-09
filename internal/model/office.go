package model

import "time"
//机构信息  
type Office struct { 
    //
    Id  int  `gorm:"column:id;PRIMARY_KEY" json:"Id,omitempty"` 
    //
    Code  string  `gorm:"column:code" json:"Code,omitempty"` 
    //
    ParentId  int  `gorm:"column:parent_id" json:"ParentId,omitempty"` 
    //
    Sname  string  `gorm:"column:sname" json:"Sname,omitempty"` 
    //
    Lname  string  `gorm:"column:lname" json:"Lname,omitempty"` 
    //
    CreateAt  time.Time  `gorm:"column:create_at" json:"CreateAt,omitempty"` 
    //
    UpdateAt  time.Time  `gorm:"column:update_at" json:"UpdateAt,omitempty"`
}

func (u Office) TableName() string {
	return "public.sys_office"
}
