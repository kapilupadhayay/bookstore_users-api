package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "CreateUser Implement me!!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "GetUser Implement me!!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "SearchUser Implement me!!")
}
