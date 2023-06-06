package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// code为0，忽略msg； 1表示拒绝，弹出msg； -1表示未知错误，弹出msg (会打印错误日志)
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Uuid string      `json:"uuid,omitempty"`
}

func Echo(ctx *gin.Context, data interface{}, err error) {
	if ctx.IsAborted() {
		logrus.Error("ctx is aborted")
		return
	}

	var r Response
	if err == nil {
		r.Msg = "ok"
		r.Data = data
	} else {
		r.Msg = err.Error()
		r.Data = struct{}{}
		r.Uuid = uuid.NewString()
		l := logrus.WithField("uuid", r.Uuid)
		// 区分 拒绝、动作、未知错误
		switch err.(type) {
		case Msg:
			r.Code = 1
			l.Info(r.Msg)
		default:
			r.Code = -1
			l.Error(r.Msg)
		}
	}

	byteData, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}

	ctx.Abort()
	ctx.Data(http.StatusOK, "application/json", byteData)

	// 这里和中间件log配合
	ctx.Set("response", byteData)
}
