package ugin

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func Context(c *gin.Context) context.Context {
	return c.Request.Context()
}

func FullLog(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(body))
	c.Request.Body = rdr2

	buf := new(bytes.Buffer)
	buf.ReadFrom(rdr1)

	start := time.Now()
	path := c.Request.URL.Path

	c.Next()

	end := time.Now()
	latency := end.Sub(start)

	status := c.Writer.Status()
	method := c.Request.Method
	ip := c.ClientIP()
	bodyf := buf.String()

	fmt.Println("")
}
