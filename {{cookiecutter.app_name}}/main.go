package main


import (
	"github.com/TaylorGWu/gin-server/gateway"
	"{{cookiecutter.app_name}}/handler"
)

func main() {
	helloWorldFront := handler.HelloWorldHandler{

	}

	{{cookiecutter.gateway_name}} := gateway.Gateway{

	}
	{{cookiecutter.gateway_name}}.AddFronts(helloWorldFront)
	{{cookiecutter.gateway_name}}.Run()
}