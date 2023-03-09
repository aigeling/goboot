package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aigeling/goboot/frame/utils"
)

type PageInfo struct {
	PageIndex *int   `json:"PageIndex,omitempty" gorm:"-"`
	PageSize  *int   `json:"PageSize,omitempty" gorm:"-"`
	RowCount  *int64 `json:"RowCount,omitempty" gorm:"-"`
	PageCount *int   `json:"PageCount,omitempty" gorm:"-"`
}

func (o *PageInfo) GetPageInfo() *PageInfo {

	pageIndex := o.GetPageIndex()
	pageSize := o.GetPageSize()
	return &PageInfo{
		PageIndex: &pageIndex,
		PageSize:  &pageSize,
		RowCount:  o.RowCount,
		PageCount: o.PageCount,
	}
}

func (o *PageInfo) SetPageIndex(pageIndex int) {
	o.PageIndex = &pageIndex
}
func (o *PageInfo) SetPageSize(pageSize int) {
	o.PageSize = &pageSize
}
func (o *PageInfo) SetRowCount(rowCount int64) *PageInfo {

	pageSize := o.GetPageSize()

	if pageSize == 0 {
		pageSize = 10
		o.SetPageSize(pageSize)
	}

	if rowCount%int64(pageSize) == 0 {
		o.SetPageCount(int(rowCount) / pageSize)
	} else {
		o.SetPageCount(int(rowCount)/pageSize + 1)
	}

	o.RowCount = &rowCount
	return o
}
func (o *PageInfo) SetPageCount(pageCount int) {
	o.PageCount = &pageCount
}

func (o *PageInfo) GetPageIndex() int {
	if o.PageIndex == nil {
		return 1
	}
	return *o.PageIndex
}
func (o *PageInfo) GetPageSize() int {
	if o.PageSize == nil {
		return 10
	}
	return *o.PageSize
}
func (o *PageInfo) GetRowCount() int64 {
	if o.RowCount == nil {
		return 0
	}
	return *o.RowCount
}
func (o *PageInfo) GetPageCount() int {
	if o.PageCount == nil {
		return 0
	}
	return *o.PageCount
}

// 获取分页offset和limit
func (o *PageInfo) GetOffSetAndLimit() (int, int) {

	pageSize := o.GetPageSize()
	pageIndex := o.GetPageIndex()

	if pageSize <= 0 {
		pageSize = 10
		o.SetPageSize(pageSize)
	}
	if pageIndex <= 0 {
		pageIndex = 1
		o.SetPageIndex(pageIndex)
	}

	offset := (pageIndex - 1) * pageSize

	return offset, pageSize

}

type PageResult[T any] struct {
	*PageInfo
	List []map[string]any `json:"list"`
}

// http返回对象
type RespVo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 创建响应结果
func NewResp(code int, msg string, data any) *RespVo {
	var rstData any

	if utils.IsSlice(data) {
		rstData = data
	} else if !utils.IsNil(data) {
		var list []any
		list = append(list, data)
		rstData = list
	}
	return &RespVo{
		Code: code,
		Msg:  msg,
		Data: rstData,
	}
}

// 响应成功
func Resp(c *gin.Context, code int, msg string, data any) {
	c.JSON(http.StatusOK, NewResp(code, msg, data))
}

// 响应成功
func RespSucc(c *gin.Context, msg string) {
	Resp(c, 0, msg, nil)
}

// 返回响应数据
func RespData(c *gin.Context, data any) {
	Resp(c, 0, "", data)
}

// 响应失败
func RespErr(c *gin.Context, code int, msg string) {
	Resp(c, code, msg, nil)
}

// 返回成功,用于更新成功返回
func NewSuccResp() *RespVo {
	return NewResp(0, "", nil)
}

// 返回数据响应,用于返回查询数据
func NewDataResp(data any) *RespVo {
	return NewResp(0, "", data)
}

// 创建失败响应
func NewErrResp(code int, msg string) *RespVo {
	return NewResp(0, msg, nil)
}
