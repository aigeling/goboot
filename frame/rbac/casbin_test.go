package rbac

import (
	"fmt"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

func TestRule(t *testing.T) {

	text := `
		[request_definition]
		r = sub, obj, act
	  
		[policy_definition]
		p = sub, obj, act
	  
		[policy_effect]
		e = some(where (p.eft == allow))
	  
		[matchers]
		m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
		`

	m, err := model.NewModelFromString(text)
	if err != nil {
		t.Errorf("err:%v", err)
	}

	e, err := casbin.NewEnforcer(m, "./policy.csv")
	if err != nil {
		t.Errorf("err:%v", err)
	}

	check(e, "zxp", "data1", "read")
	check(e, "zhang", "data2", "write")
	check(e, "zxp", "data1", "write")
	check(e, "zxp", "data2", "read")
}

func check(e *casbin.Enforcer, sub, obj, act string) {

	ok, _ := e.Enforce(sub, obj, act)

	// if err != nil {
	// 	return false
	// }
	// return ok
	if ok {
		fmt.Printf("%s CAN %s %s  \n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func TestPermision(t *testing.T) {
	InitCasbin()

	// AddPolicy("admin", "domain1", "/a/hello/:id", "GET")

	// AddRoleForUser("zhangsan", "admin", "domain1")
	// enforcer.AddPolicy("zhangsan", "admin", "domain1")

	rst := CheckPermission("zhangsan", "domain1", "/a/sys/office/add", "POST")
	t.Logf("check office :%v", rst)
	rst = CheckPermission("zhangsan", "domain1", "/a/sys/user/add", "POST")
	t.Logf("check user :%v", rst)
}

func TestGetUserPermissions(t *testing.T) {
	InitCasbin()

	AddPolicy("admin", "domain1", "/a/hello/:id", "GET")
	AddPolicy("admin", "domain1", "/a/hello1/:id", "GET")
	AddPolicy("worker", "domain1", "/a/hello1/:id", "GET")
	AddRoleForUser("zhangsan", "admin", "domain1")
	AddRoleForUser("zhangsan", "worker", "domain1")
	rst, _ := enforcer.GetImplicitPermissionsForUser("zhangsan", "domain1")

	t.Logf("list: %v ", rst)

}
