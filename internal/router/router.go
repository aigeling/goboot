package router

import (
	"errors"

	"github.com/aigeling/goboot/frame/rbac"
	"github.com/aigeling/goboot/frame/web"
	"github.com/aigeling/goboot/global"
	"github.com/aigeling/goboot/internal/controller"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouteRegister interface {
	RegitsRouter(r *gin.RouterGroup)
}

// 初始化路由
func InitRouter(g *gin.Engine) {

	authGroup := g.Group("/a")

	//挂载jwt
	jwtMaker := &web.JwtAuthMiddelWareMaker{
		Engine:               g,
		AuthGroups:           []*gin.RouterGroup{authGroup},
		AuthorizationHandler: authorition,
		LoginHandler:         login,
		SecretKey:            global.Conf.Server.JwtSecretKey,
	}
	jwtMaker.MountJwt()

	initUnAuthGroup(g)
	initAuthGroup(authGroup)
}

// 不需要认证的router
func initUnAuthGroup(g *gin.Engine) {
	g.GET("hello", func(c *gin.Context) {
		c.JSON(200, "hello ok")
	})
}

func addController(gg *gin.RouterGroup, rr RouteRegister) {
	rr.RegitsRouter(gg)
}

// 初始需要认证的router
func initAuthGroup(gg *gin.RouterGroup) {

	gg.POST("gen", web.GenFile)
	sysGroup := gg.Group("sys")
	addController(sysGroup, controller.User)
	addController(sysGroup, controller.Office)
}

type User struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

func authorition(loginName string, c *gin.Context) bool {

	// fmt.Printf("auth loginName:%s , uri:%s\n", loginName, c.Request.RequestURI)

	rst := rbac.CheckPermission(loginName, "domain1", c.Request.RequestURI, c.Request.Method)
	zap.L().Sugar().Debugf("%s %s %s --> %v", loginName, c.Request.Method, c.Request.RequestURI, rst)

	return rst
}

func login(c *gin.Context) (loginName string, err error) {

	user := &User{}
	c.ShouldBind(user)
	if user.LoginName == "zhangsan" && user.Password == "123456" {
		return user.LoginName, nil
	}
	if user.LoginName == "sysadmin" && user.Password == "123456" {
		return user.LoginName, nil
	}
	if user.LoginName == "lisi" && user.Password == "123456" {
		return user.LoginName, nil
	}
	return "", errors.New("login failed")
}
