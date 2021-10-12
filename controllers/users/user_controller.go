package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jademnp/go-store-user-api/domain/users"
	"github.com/jademnp/go-store-user-api/services"
	"github.com/jademnp/go-store-user-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.BadRequestError("bad json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.BadRequestError("user id should be number")
		c.JSON(err.Status, err)
		return
	}
	result, saveErr := services.GetUser(userId)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
