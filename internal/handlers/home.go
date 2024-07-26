package handlers

import (
	"api/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	_, logged := c.Get("user")

	elem := views.Base(logged)

	templ.Handler(elem).ServeHTTP(c.Writer, c.Request)
}
