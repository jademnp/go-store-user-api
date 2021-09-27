package app

import "github.com/jademnp/go-store-user-api/controllers/users"

func url_mapping() {
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
