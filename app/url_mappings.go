package app

import (
	"github.com/kapilupadhayay/bookstore_users-api/controllers/ping"
	"github.com/kapilupadhayay/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
}
