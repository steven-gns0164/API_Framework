package main

import (
	"gin_framework/controllers"

	"github.com/gin-gonic/gin"
)

// go get -u github.com/gin-gonic/gin
func main() {
	r := gin.Default()

	r.GET("/controller", controllers.GetAllStudents)
	r.POST("/controller", controllers.InsertNewStudents)
	r.PUT("/controller", controllers.UpdateStudent)
	r.DELETE("/controller", controllers.DeleteStudent)

	r.Run()
}
