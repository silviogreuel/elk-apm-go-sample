package handlers

import (
	"github.com/eventials/vlab-boilerplate-api/pkg/handlers/hello"
	"github.com/gin-gonic/gin"
)

func Serve(addr string) error {
	api := gin.New()
	hello.Attach(api.Group("/hello"))
	return api.Run(addr)
}
