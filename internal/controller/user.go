package controller

import (
	"github.com/aigeling/goboot/frame/web"
	"github.com/aigeling/goboot/internal/entity"
	"github.com/aigeling/goboot/internal/service"
	"github.com/gin-gonic/gin"
)

type userController struct{}

var User = &userController{}

func (o *userController) RegitsRouter(r *gin.RouterGroup) {

	r.POST("/user/add", User.Add)
	r.POST("/user/update", User.Update)
	r.POST("/user/get", User.Get)
	r.POST("/user/delete", User.Delete)
	r.POST("/user/list", User.List)
	r.POST("/user/page", User.Page)
}

func (o *userController) Add(c *gin.Context) {

	user := &entity.User{}
	err := binder.BindBodyExcept(c, user, "Id")
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	err = service.UserService.Insert(user)
	if err != nil {
		logger.Errorf("Add user error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespSucc(c, "新增用户信息成功")
}

func (o *userController) List(c *gin.Context) {
	user := &entity.User{}
	err := binder.BindBody(c, user)
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, err := service.UserService.QueryList(user)
	if err != nil {
		logger.Errorf("get user error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespData(c, resp)
}

func (o *userController) Page(c *gin.Context) {
	user := &entity.User{}
	err := binder.BindBody(c, user)
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, total, err := service.UserService.FindPage(user)
	if err != nil {
		logger.Errorf("get user error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespData(c, web.PageResult[entity.User]{
		List:     resp,
		PageInfo: user.GetPageInfo().SetRowCount(total),
	})

}

func (o *userController) Get(c *gin.Context) {

	user := &entity.User{}
	err := binder.BindBody(c, user, "Id")
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, err := service.UserService.GetByPk(user)
	if err != nil {
		logger.Errorf("get user error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}

	web.RespData(c, resp)
}

func (o *userController) Update(c *gin.Context) {

	user := &entity.User{}
	err := binder.BindBodyExcept(c, user)
	if err != nil {
		logger.Errorf("user parameter error:%v", err)
		web.RespErr(c, 501, err.Error())
		return
	}

	err = service.UserService.UpdateByPk(user)
	if err != nil {
		logger.Errorf("update user error:%v", err)
		web.RespErr(c, 500, err.Error())
		return
	}

	web.RespSucc(c, "修改用户信息成功")

}

func (o *userController) Delete(c *gin.Context) {

	user := &entity.User{}
	err := binder.BindBody(c, user, "Id")
	if err != nil {
		logger.Errorf("user parameter error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	err = service.UserService.DeleteByPk(user)
	if err != nil {
		logger.Errorf("delete user error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}

	web.RespSucc(c, "删除用户信息成功")

}
