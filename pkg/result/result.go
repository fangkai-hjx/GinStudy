package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Ctx *gin.Context
}

type ResultCont struct {
	Code int         `json:"code"` //提示代码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"`
}

func NewResult(c *gin.Context) *Result {
	return &Result{Ctx: c}
}

//成功
func (r *Result) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultCont{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	r.Ctx.JSON(http.StatusOK, res)
}

//出错
func (r *Result) Error(code int, msg string) {
	res := ResultCont{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	}
	r.Ctx.JSON(http.StatusOK, res)
}
