package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_SESSION_NAME = "GSESSION"
)

func SetSession(c *gin.Context, name, value any) error {
	s := sessions.Default(c)
	s.Set(name, value)
	return s.Save()
}

func DelSession(c *gin.Context, name any) error {
	s := sessions.Default(c)
	s.Delete(name)
	return s.Save()
}

func GetSession(c *gin.Context, name any) any {
	s := sessions.Default(c)
	return s.Get(name)
}

func GetSessionStr(c *gin.Context, name any) string {
	s := sessions.Default(c)
	obj := s.Get(name)
	if obj == nil {
		return ""
	} else {
		return obj.(string)
	}
}

func ClearSession(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
}

// 初始化session
func InitSession(ge *gin.Engine, sessionMaxAge int, sessionStore sessions.Store) {

	var store sessions.Store
	if sessionStore != nil {
		store = sessionStore
	} else {
		store = memstore.NewStore([]byte("MemSession"))
	}
	if sessionMaxAge == 0 {
		sessionMaxAge = 1800
	}
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
	})
	ge.Use(sessions.Sessions(DEFAULT_SESSION_NAME, store))
}
