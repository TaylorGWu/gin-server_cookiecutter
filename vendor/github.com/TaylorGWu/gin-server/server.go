package main

import (
	"gin-server/gateway"
	"gin-server/protocol"
	"github.com/gin-gonic/gin"
)

type Front1 struct {

}

func (front Front1) BindMethod() (string) {
	return protocol.HTTP_GET_METHOD
}

func (front Front1) BindPath() (string) {
	return "/test1"
}

func (front Front1) Serve(context *gin.Context) {
	context.Writer.Write([]byte("hello,world"))

}

func main() {
	front1 := Front1{}
	myGateway := gateway.Gateway{
	}
	myGateway.AddFronts(front1)
	myGateway.Run()
}
