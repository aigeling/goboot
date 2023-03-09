package entity

import (
	"reflect"
	"strings"
	"time"

	"github.com/aigeling/goboot/frame/web"
	"github.com/aigeling/goboot/internal/model"
)

// 用户信息
type User struct {
	web.PageInfo
	//
	Id *int64 `json:"Id,omitempty" validate:"required"`
	//
	Code *string `json:"code,omitempty" `
	//用户登录名
	LoginName *string `json:"LoginName,omitempty" validate:"required"`
	//密码
	Password *string `json:"Password,omitempty" `
	//用户姓名
	Name *string `json:"Name,omitempty" validate:"required"`
	//
	GroupId *int64 `json:"GroupId,omitempty" `
	//
	CreatedAt *time.Time `json:"CreatedAt,omitempty" `
	//
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty" `
}

// 转换为map格式，以tag中的json name为key，若json name为空则为字段名
func (o *User) ToRespMap() map[string]any {

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
func (o *User) ToMap() map[string]any {
	m := make(map[string]any)
	if o.Id != nil {
		m["Id"] = *o.Id
	}
	if o.Code != nil {
		m["Code"] = *o.Code
	}
	if o.LoginName != nil {
		m["LoginName"] = *o.LoginName
	}
	if o.Password != nil {
		m["Password"] = *o.Password
	}
	if o.Name != nil {
		m["Name"] = *o.Name
	}
	if o.GroupId != nil {
		m["GroupId"] = *o.GroupId
	}
	if o.CreatedAt != nil {
		m["CreatedAt"] = *o.CreatedAt
	}
	if o.UpdatedAt != nil {
		m["UpdatedAt"] = *o.UpdatedAt
	}

	return m
}

// 转为model格式
func (o *User) ToModel() *model.User {
	m := &model.User{}
	if o.Id != nil {
		m.Id = *o.Id
	}
	if o.Code != nil {
		m.Code = *o.Code
	}
	if o.LoginName != nil {
		m.LoginName = *o.LoginName
	}
	if o.Password != nil {
		m.Password = *o.Password
	}
	if o.Name != nil {
		m.Name = *o.Name
	}
	if o.GroupId != nil {
		m.GroupId = *o.GroupId
	}
	if o.CreatedAt != nil {
		m.CreatedAt = *o.CreatedAt
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = *o.UpdatedAt
	}

	return m

}

func (o *User) GetId() int64 {
	return *o.Id
}

func (o *User) GetCode() string {
	return *o.Code
}

func (o *User) GetLoginName() string {
	return *o.LoginName
}

func (o *User) GetPassword() string {
	return *o.Password
}

func (o *User) GetName() string {
	return *o.Name
}

func (o *User) GetGroupId() int64 {
	return *o.GroupId
}

func (o *User) GetCreatedAt() time.Time {
	return *o.CreatedAt
}

func (o *User) GetUpdatedAt() time.Time {
	return *o.UpdatedAt
}

func (o *User) SetId(id int64) {
	o.Id = &id
}

func (o *User) SetCode(code string) {
	o.Code = &code
}

func (o *User) SetLoginName(loginName string) {
	o.LoginName = &loginName
}

func (o *User) SetPassword(password string) {
	o.Password = &password
}

func (o *User) SetName(name string) {
	o.Name = &name
}

func (o *User) SetGroupId(groupId int64) {
	o.GroupId = &groupId
}

func (o *User) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = &createdAt
}

func (o *User) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = &updatedAt
}
