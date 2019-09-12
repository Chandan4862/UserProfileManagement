package main

import (
	"net/http"
	api "project/GoTrainingAssignment/UserProfileManagement/api"
	helper "project/GoTrainingAssignment/UserProfileManagement/api/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	helper.InitDb()
	startServer()
}
func startServer() {
	route := gin.Default()
	api.Init(route)
	route.GET("/", check())
	s	 := &http.Server{
		Addr:    ":8080",
		Handler: route,
	}
	s.ListenAndServe()
}
func check() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}
