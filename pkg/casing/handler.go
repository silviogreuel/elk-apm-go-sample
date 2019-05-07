package casing

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silviogreuel/elk-apm-go-sample/pkg/utils/uapm"
	"github.com/silviogreuel/elk-apm-go-sample/pkg/utils/ugin"
)

func Attach(g *gin.RouterGroup) {
	g.GET("/upper/:value", func(c *gin.Context) {
		ctx := ugin.Context(c)
		defer uapm.StartSpan(ctx, "casing.Attach", "handler").End()

		value := c.Param("value")

		result := ToUpper(ctx, value)

		c.String(http.StatusOK, result)
	})

	g.GET("/lower/:value", func(c *gin.Context) {
		ctx := ugin.Context(c)
		defer uapm.StartSpan(ctx, "casing.Attach", "handler").End()

		value := c.Param("value")

		result := ToLower(ctx, value)

		c.String(http.StatusOK, result)
	})

	g.GET("/case/:value", func(c *gin.Context) {
		ctx := ugin.Context(c)
		defer uapm.StartSpan(ctx, "casing.Attach", "handler").End()

		casing := c.Query("casing")
		value := c.Param("value")

		result, err := CaseValue(ctx, casing, value)

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.String(http.StatusOK, result)
		}
	})
}
