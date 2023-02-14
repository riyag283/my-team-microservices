package main

import (
	"authentication/app/controllers"
	"authentication/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitialiseDBConnection()

	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	if err := r.Run("localhost:9090"); err != nil {
		panic(err.Error())
	}
}