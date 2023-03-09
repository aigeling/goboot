package web

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	identityKey        = "id"
	payload_login_name = "login_name"
)

// jwt认证中间件
type JwtAuthMiddelWareMaker struct {
	Engine *gin.Engine
	//需要进行认证的路由分组,认证中间件将作用在该分组上
	AuthGroups []*gin.RouterGroup
	//认证中间件实例
	authMiddleware *jwt.GinJWTMiddleware
	//自定义登录处理,登录成功后返回登录用户名
	LoginHandler func(c *gin.Context) (loginName string, err error)
	//自定义权限认证,根据登录用户名,验证其是否有权限访问，返回是否成功
	AuthorizationHandler func(loginName string, c *gin.Context) bool

	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	// - "param:<name>"
	// TokenLookup: "header: Authorization, query: token, cookie: jwt",
	TokenLookup string

	//登录提交url (post)，默认为：/login
	LoginUrl string

	//jwt 加密密钥
	SecretKey string
}

// 将jwt中间件挂载到gin上
func (o *JwtAuthMiddelWareMaker) MountJwt() (err error) {
	err = o.createAuthMiddleware()
	if err != nil {
		return
	}
	o.Engine.GET("/refresh_token", o.authMiddleware.RefreshHandler)
	if o.LoginUrl == "" {
		o.LoginUrl = "/login"
	}
	o.Engine.POST(o.LoginUrl, o.authMiddleware.LoginHandler)
	// o.Engine.NoRoute(o.authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
	o.Engine.NoRoute(func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "msg": "Page not found"})
	})

	for _, authGroup := range o.AuthGroups {
		authGroup.Use(o.authMiddleware.MiddlewareFunc())
	}
	return
}

// 创建认证中间件
func (o *JwtAuthMiddelWareMaker) createAuthMiddleware() (err error) {

	if o.SecretKey == "" {
		o.SecretKey = "I'm JWT Auth Middel Ware"
	}

	// the jwt middleware
	o.authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "jwt-realm",
		Key:         []byte(o.SecretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			loginName, err := o.LoginHandler(c)

			if err != nil {
				log.Printf("login auth fail:%v", err)
				return nil, jwt.ErrFailedAuthentication
			}

			return loginName, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {

			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					payload_login_name: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			loginName := claims[payload_login_name].(string)

			c.Set(payload_login_name, loginName)
			return loginName
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {

			if loginName, ok := data.(string); ok {
				return o.AuthorizationHandler(loginName, c)
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code": code,
				"msg":  message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		// TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "header: Authorization",
		TokenLookup: o.tokenLookup(),
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return
	}
	return
}

func (o *JwtAuthMiddelWareMaker) tokenLookup() string {

	if o.TokenLookup == "" {
		return "header: Authorization"
	} else {
		return o.TokenLookup
	}
}

func GetLoginName(c *gin.Context) string {
	return c.GetString(payload_login_name)
}
