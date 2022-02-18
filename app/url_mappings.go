package app

import (
	"users-api/controllers/ping_controller"
	"users-api/controllers/users_controller"
)

func MapUrls() {
	router.GET("/ping", ping_controller.Ping)
	router.GET("/users/:user_id", users_controller.GetUser)
	router.POST("/users", users_controller.CreateUser)
}
