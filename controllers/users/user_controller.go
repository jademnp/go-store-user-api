package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "ok ja")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "ok ja")
}
