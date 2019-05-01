package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Attach(g *gin.RouterGroup) {
	g.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
}
