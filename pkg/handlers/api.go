package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/silviogreuel/elk-apm-go-sample/pkg/casing"
	"go.elastic.co/apm/module/apmgin"
)

func Serve(addr string) error {
	api := gin.New()
	api.Use(apmgin.Middleware(api))
	casing.Attach(api.Group("/casing"))
	return api.Run(addr)
}
