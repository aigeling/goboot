package rbac

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// var permissionMap map[string]*resource = make(map[string]*resource)
var enforcer *casbin.Enforcer

type HandlerRegister interface {
	Handle(httpMethod string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	BasePath() string
}

// type resource struct {
// 	URI        string `json:"uri"`
// 	Permission string `json:"permission"`
// 	Name       string `json:"name"`
// }

// 注册路由，通过该函数可以增加uri跟权限标识的对应关系。
// HandlerRegister 可以是*gin.Engine或*gin.RouterGroup
// func RegisteRouter(g HandlerRegister, method, uri, name, permission string, handlers ...gin.HandlerFunc) {

// 	u := path.Join(g.BasePath(), uri)
// 	permissionMap[u] = &resource{
// 		URI:        uri,
// 		Name:       name,
// 		Permission: permission,
// 	}

// 	g.Handle(method, uri, handlers...)
// }

// func RegisteGet(g HandlerRegister, uri, name, permission string, handlers ...gin.HandlerFunc) {
// 	RegisteRouter(g, METHOD_GET, uri, name, permission, handlers...)
// }
// func RegistePost(g HandlerRegister, uri, name, permission string, handlers ...gin.HandlerFunc) {
// 	RegisteRouter(g, METHOD_POST, uri, name, permission, handlers...)
// }
// func RegistePut(g HandlerRegister, uri, name, permission string, handlers ...gin.HandlerFunc) {
// 	RegisteRouter(g, METHOD_PUT, uri, name, permission, handlers...)
// }
// func RegisteDelete(g HandlerRegister, uri, name, permission string, handlers ...gin.HandlerFunc) {
// 	RegisteRouter(g, METHOD_DELETE, uri, name, permission, handlers...)
// }

// 获取登录用户所有权限
func GetPermissions(user string) ([][]string, error) {

	return enforcer.GetImplicitPermissionsForUser(user, "*")

}

// 初始化casbin
func InitCasbin() error {
	// text := `
	// [request_definition]
	// r = sub, dom, obj, act

	// [policy_definition]
	// p = sub, dom, obj, act

	// [role_definition]
	// g = _, _, _

	// [policy_effect]
	// e = some(where (p.eft == allow))

	// [matchers]
	// m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && keyMatch2(r.obj, p.obj) &&  r.act == p.act
	// `

	// m, err := model.NewModelFromString(text)
	// if err != nil {
	// 	return err
	// }
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	// e, err := casbin.NewEnforcer("D:\\work-golang\\goboot\\model.conf", "D:\\work-golang\\goboot\\policy.csv")
	// e, err := casbin.NewEnforcer(m)
	if err != nil {
		return err
	}

	enforcer = e

	return nil
}

// 检查用户是否有权限
func CheckPermission(sub, domain, obj, act string) bool {

	ok, err := enforcer.Enforce(sub, domain, obj, act)

	if err != nil {
		logger.Errorf("check permission:%v", err)
		return false
	}
	return ok
}

// 添加策略
// role 角色；
// domain 所属域，可以类比机构，比如统一角色可以在不同法人机构有不同权限；
// obj 访问对象，一般设置为uri，支持rest api ;
// act 对访问对象的操作，一般为访问uri的method，如GET，POST ;
func AddPolicy(role, domain, obj, act string) (bool, error) {
	return enforcer.AddPolicy(role, domain, obj, act)
}

// 给用户添加角色
func AddRoleForUser(user, role string, domains ...string) (bool, error) {
	return enforcer.AddRoleForUser(user, role, domains...)
}

// func initPolicy() {
// 	AddPolicy("admin", "domain1", "/a/permissions", METHOD_GET)
// 	AddPolicy("admin", "domain1", "/a/hello/:id", METHOD_GET)
// 	AddRoleForUser("zhangsan", "admin", "domain1")
// }
