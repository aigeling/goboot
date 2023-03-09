package entity

import (
	"testing"
	"time"
)

type A struct {
	Num *int
}

func (o *A) Add1() {

	n := 1
	if o.Num == nil {
		o.Num = &n
	}

	n = *(o.Num) + 1

	o.Num = &n
}

func TestAaa(t *testing.T) {

	a := &A{}
	a.Add1()
	t.Log("2:", *a.Num)

}

func TestMap(t *testing.T) {

	u := &User{}
	u.SetCode("123")
	u.SetId(1)
	u.SetName("zhangsan")
	u.SetCreatedAt(time.Now())

	// u.SetPageIndex(1)
	// u.SetPageSize(10)
	// t.Log(utils.TagJsonName(*u, "Name"))

	// t.Log(u.ToRespMap())
	t.Log(u.ToMap())
}
