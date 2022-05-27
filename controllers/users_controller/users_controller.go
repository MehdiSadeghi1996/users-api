package users_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"users-api/domain/users"
	"users-api/services"
	"users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
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
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be numeric")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

func UpdateUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be numeric")
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId

	isPartialOrPatch := c.Request.Method == http.MethodPatch
	result, err := services.UpdateUser(isPartialOrPatch, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)

}
