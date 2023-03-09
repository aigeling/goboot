package service

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/aigeling/goboot/frame/config"
	"github.com/aigeling/goboot/frame/db"
	"github.com/aigeling/goboot/global"
	"github.com/aigeling/goboot/internal/entity"
)

func TestSelect(t *testing.T) {

	config.InitConfig(global.Conf, "app")
	db.InitDB()
	t.Logf("defaultDB :  %s", global.Conf.DefaultDB)
	db := db.GetDefaultDB()

	if db == nil {
		t.Fatal("db config err")
	}
	serv := NewUserService(db)

	user, err := serv.GetByPk(&entity.User{Id: Int64(2)})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("user: %+v", user)
	// t.Logf("user: %+v", user.ToMap())

}

func Int64(i int64) *int64 {
	return &i
}

func Pointer[T comparable](t T) *T {
	return &t
}

func ToMap(content interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		return nil, err
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&m); err != nil {
			return nil, err
		} else {
			for k, v := range m {
				m[k] = v
			}
		}
	}
	return m, nil
}

type ABC struct {
	name string
}

func TestAppend(t *testing.T) {

	// v := reflect.ValueOf(l)
	// t.Log("kind", v.Kind())
}

func IsSlice[T any](o any) bool {

	_, ok := o.([]T)

	return ok

}
