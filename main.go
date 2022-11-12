package main

import (
	"github.com/GabrielEdwinSP/GolangDeveloperTest/controllers"
	"github.com/GabrielEdwinSP/GolangDeveloperTest/initializers"
	"github.com/GabrielEdwinSP/GolangDeveloperTest/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	//API No 1
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	//API No 2
	r.POST("/joblist", controllers.RequestJobList)
	r.GET("/joblist", controllers.RequestJobListSearch)

	r.Run()
}
