package hello

import (
	"net/http"

	domain "github.com/eventials/vlab-boilerplate-api/pkg/hello"

	"github.com/gin-gonic/gin"
)

func Attach(g *gin.RouterGroup) {
	g.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		result := domain.SayHello(name)

		c.String(http.StatusOK, result)
	})
}
