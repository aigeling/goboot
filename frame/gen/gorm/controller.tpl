package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/aigeling/goboot/frame/web"
	"{{.PackageBasePath}}/entity"
	"{{.PackageBasePath}}/service"
)
{{$objectName := LowerFirstChar .ObjectName}}
type {{$objectName}}Controller struct{}

var {{.ObjectName}} = &{{$objectName}}Controller{}

func (o *{{$objectName}}Controller) RegitsRouter(r *gin.RouterGroup) {

	r.POST("/{{$objectName}}/add", {{.ObjectName}}.Add)
	r.POST("/{{$objectName}}/update", {{.ObjectName}}.Update)
	r.POST("/{{$objectName}}/get", {{.ObjectName}}.Get)
	r.POST("/{{$objectName}}/delete", {{.ObjectName}}.Delete)
	r.POST("/{{$objectName}}/list", {{.ObjectName}}.List)
	r.POST("/{{$objectName}}/page", {{.ObjectName}}.Page)
}

func (o *{{$objectName}}Controller) Add(c *gin.Context) {

	{{$objectName}} := &entity.{{.ObjectName}}{}
	err := binder.BindBodyExcept(c, {{$objectName}}, "Id")
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	err = service.{{.ObjectName}}Service.Insert({{$objectName}})
	if err != nil {
		logger.Errorf("Add {{$objectName}} error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespSucc(c, "新增{{.Remark}}成功")
}

func (o *{{$objectName}}Controller) List(c *gin.Context) {
	{{$objectName}} := &entity.{{.ObjectName}}{}
	err := binder.BindBody(c, {{$objectName}})
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, err := service.{{.ObjectName}}Service.QueryList({{$objectName}})
	if err != nil {
		logger.Errorf("get {{$objectName}} error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespData(c, resp)
}

func (o *{{$objectName}}Controller) Page(c *gin.Context) {
	{{$objectName}} := &entity.{{.ObjectName}}{}
	err := binder.BindBody(c, {{$objectName}})
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, total, err := service.{{.ObjectName}}Service.FindPage({{$objectName}})
	if err != nil {
		logger.Errorf("get {{$objectName}} error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespData(c, web.PageResult[entity.{{.ObjectName}}]{
		List:     resp,
		PageInfo: {{$objectName}}.GetPageInfo().SetRowCount(total),
	})

}

func (o *{{$objectName}}Controller) Get(c *gin.Context) {

	{{$objectName}} := &entity.{{.ObjectName}}{}
	err := binder.BindBody(c, {{$objectName}}, "Id")
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, err := service.{{.ObjectName}}Service.GetByPk({{$objectName}})
	if err != nil {
		logger.Errorf("get {{$objectName}} error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}

	web.RespData(c, resp)
}

func (o *{{$objectName}}Controller) Update(c *gin.Context) {

	{{$objectName}} := &entity.{{.ObjectName}}{}
	err := binder.BindBodyExcept(c, {{$objectName}})
	if err != nil {
		logger.Errorf("{{$objectName}} parameter error:%v", err)
		web.RespErr(c, 501, err.Error())
		return
	}

	err = service.{{.ObjectName}}Service.UpdateByPk({{$objectName}})
	if err != nil {
		logger.Errorf("update {{$objectName}} error:%v", err)
		web.RespErr(c, 500, err.Error())
		return
	}

	web.RespSucc(c, "修改{{.Remark}}成功")

}

func (o *{{$objectName}}Controller) Delete(c *gin.Context) {

	{{$objectName}} := &entity.{{.ObjectName}}{}
	err := binder.BindBody(c, {{$objectName}}, "Id")
	if err != nil {
		logger.Errorf("{{$objectName}} parameter error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	err = service.{{.ObjectName}}Service.DeleteByPk({{$objectName}})
	if err != nil {
		logger.Errorf("delete {{$objectName}} error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}

	web.RespSucc(c, "删除用户信息成功")

}
