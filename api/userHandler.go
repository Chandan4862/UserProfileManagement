package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewUser()
		c.Bind(&requestBody)
		AddUserService(requestBody)
		c.JSON(http.StatusOK, GetUserList())
	}
}

func ViewUser() gin.HandlerFunc {
	GetAllUser()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, GetUserList())
	}
}


func DeleteUser() gin.HandlerFunc{
	return func(c *gin.Context){
		requestBody:= c.Param("id")
		c.Bind(&requestBody)
		fmt.Println("*****",requestBody)
		DeleteService(requestBody)
		c.JSON(http.StatusOK, GetUserList())
	}
}