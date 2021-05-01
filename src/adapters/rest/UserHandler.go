package rest

import (
	"github.com/gin-gonic/gin"
)

func GetUserMiddleware(c *gin.Context) {
	c.Next()
}