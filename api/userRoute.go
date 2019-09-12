package api

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/view",ViewUser())
	r.POST("/add", AddUser())
	r.DELETE("/delete/:id", DeleteUser())
}
