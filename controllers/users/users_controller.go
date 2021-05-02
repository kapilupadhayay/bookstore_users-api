package users

import (
	"net/http"

	"github.com/federicoleon/bookstore_oauth-api/src/domain/users"
	"github.com/gin-gonic/gin"
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
	c.String(http.StatusNotImplemented, "GetUser Implement me!!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "SearchUser Implement me!!")
}
