package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Error 失败处理
func Error(c *gin.Context, code int, msg string) {
	var res Response
	res.Msg = msg
	//logger.Error(res.Msg)
	c.JSON(http.StatusOK, res.ReturnError(code))
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageNo = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}
func PageError(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, code int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageNo = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnError(code))
}

func Paginate(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pageNo
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// Custom 兼容函数 .
func Custom(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}
