package web

import (
	"github.com/aigeling/goboot/frame/gen"
	"github.com/aigeling/goboot/frame/utils"
	"github.com/gin-gonic/gin"
)

var binder *utils.GinBinder = utils.NewGinBinder()

func GenFile(c *gin.Context) {

	t := &gen.TableTemplate{}

	err := binder.BindBody(c, t)

	if err != nil {
		RespErr(c, 1, err.Error())
		return
	}

	mq, err := gen.GetMetaQueryer(t.DBCode)

	if err != nil {
		RespErr(c, 2, err.Error())
		return
	}
	err = gen.GenerateFile(t, mq)
	if err != nil {
		RespErr(c, 3, err.Error())
		return
	}
	logger.Infof("Generate Code OK , table_name:%s", t.TableName)

	RespSucc(c, "Generate Code OK")
}
