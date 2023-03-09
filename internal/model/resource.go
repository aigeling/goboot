package model

import "time"
//资源信息  
type Resource struct { 
    //
    Id  int64  `gorm:"column:id;PRIMARY_KEY" json:"Id,omitempty"` 
    //资源代码
    Code  string  `gorm:"column:code" json:"Code,omitempty"` 
    //资源名称
    Name  string  `gorm:"column:name" json:"Name,omitempty"` 
    //资源对象，一般指uri
    Obj  string  `gorm:"column:obj;PRIMARY_KEY" json:"Obj,omitempty"` 
    //执行操作，一般指http请求method
    Act  string  `gorm:"column:act;PRIMARY_KEY" json:"Act,omitempty"` 
    //
    CreatedAt  time.Time  `gorm:"column:created_at" json:"CreatedAt,omitempty"` 
    //
    UpdatedAt  time.Time  `gorm:"column:updated_at" json:"UpdatedAt,omitempty"`
}

func (u Resource) TableName() string {
	return "public.sys_resource"
}
