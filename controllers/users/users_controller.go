package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kapilupadhayay/bookstore_users-api/domain/users"
	"github.com/kapilupadhayay/bookstore_users-api/services"
	"github.com/kapilupadhayay/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restError.Status, restError)
		return
	}
	result, createError := services.CreateUser(user)
	if createError != nil {
		c.JSON(createError.Status, createError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restError := errors.NewBadRequestError("Invalid User Id")
		c.JSON(restError.Status, restError)
		return
	}
	result, getError := services.GetUser(user_id)
	if getError != nil {
		c.JSON(getError.Status, getError)
	}
	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "SearchUser Implement me!!")
}
