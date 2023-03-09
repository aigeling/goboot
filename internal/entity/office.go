package entity

import (
	"reflect"
	"strings"
	"time"

	"github.com/aigeling/goboot/internal/model"

	"github.com/aigeling/goboot/frame/web"
)

// 机构信息
type Office struct {
	web.PageInfo
	//
	Id *int `json:"Id,omitempty" validate:"required"`
	//
	Code *string `json:"Code,omitempty" `
	//
	ParentId *int `json:"ParentId,omitempty" `
	//
	Sname *string `json:"Sname,omitempty" `
	//
	Lname *string `json:"Lname,omitempty" `
	//
	CreateAt *time.Time `json:"CreateAt,omitempty" `
	//
	UpdateAt *time.Time `json:"UpdateAt,omitempty" `
}

// 转换为map格式，以tag中的json name为key，若json name为空则为字段名
func (o *Office) ToRespMap() map[string]any {

	t := reflect.TypeOf(*o)
	v := reflect.ValueOf(*o)
	m := make(map[string]any, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := strings.Split(f.Tag.Get("json"), ",")[0]
		if name == "" {
			name = f.Name
		}

		if !v.Field(i).IsNil() {
			m[name] = v.Field(i).Elem().Interface()
		}
	}
	return m

}

// 转为map格式,以字段名为key
func (o *Office) ToMap() map[string]any {
	m := make(map[string]any)
	if o.Id != nil {
		m["Id"] = *o.Id
	}
	if o.Code != nil {
		m["Code"] = *o.Code
	}
	if o.ParentId != nil {
		m["ParentId"] = *o.ParentId
	}
	if o.Sname != nil {
		m["Sname"] = *o.Sname
	}
	if o.Lname != nil {
		m["Lname"] = *o.Lname
	}
	if o.CreateAt != nil {
		m["CreateAt"] = *o.CreateAt
	}
	if o.UpdateAt != nil {
		m["UpdateAt"] = *o.UpdateAt
	}

	return m
}

// 转为model格式
func (o *Office) ToModel() *model.Office {
	m := &model.Office{}
	if o.Id != nil {
		m.Id = *o.Id
	}
	if o.Code != nil {
		m.Code = *o.Code
	}
	if o.ParentId != nil {
		m.ParentId = *o.ParentId
	}
	if o.Sname != nil {
		m.Sname = *o.Sname
	}
	if o.Lname != nil {
		m.Lname = *o.Lname
	}
	if o.CreateAt != nil {
		m.CreateAt = *o.CreateAt
	}
	if o.UpdateAt != nil {
		m.UpdateAt = *o.UpdateAt
	}

	return m

}

func (o *Office) GetId() int {
	return *o.Id
}

func (o *Office) GetCode() string {
	return *o.Code
}

func (o *Office) GetParentId() int {
	return *o.ParentId
}

func (o *Office) GetSname() string {
	return *o.Sname
}

func (o *Office) GetLname() string {
	return *o.Lname
}

func (o *Office) GetCreateAt() time.Time {
	return *o.CreateAt
}

func (o *Office) GetUpdateAt() time.Time {
	return *o.UpdateAt
}

func (o *Office) SetId(id int) {
	o.Id = &id
}

func (o *Office) SetCode(code string) {
	o.Code = &code
}

func (o *Office) SetParentId(parentId int) {
	o.ParentId = &parentId
}

func (o *Office) SetSname(sname string) {
	o.Sname = &sname
}

func (o *Office) SetLname(lname string) {
	o.Lname = &lname
}

func (o *Office) SetCreateAt(createAt time.Time) {
	o.CreateAt = &createAt
}

func (o *Office) SetUpdateAt(updateAt time.Time) {
	o.UpdateAt = &updateAt
}
