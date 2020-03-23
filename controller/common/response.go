package common

import (
	"wemall-go/model"
	"github.com/kataras/iris"
)

// SendErrJSON 有错误发生时，发送错误JSON
func SendErrJSON(msg string, ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{
		"errNo" : model.ErrorCode.ERROR,
		"msg"   : msg,
		"data"  : iris.Map{},
	})
}