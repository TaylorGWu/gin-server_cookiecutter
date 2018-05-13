package handler

import (
	"gin-server/protocol"
	"github.com/gin-gonic/gin"
)

type HelloWorldHandler struct {

}

func (front HelloWorldHandler) BindMethod() (string) {
	return protocol.HTTP_GET_METHOD
}

func (front HelloWorldHandler) BindPath() (string) {
	return "/helloworld"
}

func (front HelloWorldHandler) Serve(context *gin.Context) {
	context.Writer.Write([]byte("hello,world"))
}
