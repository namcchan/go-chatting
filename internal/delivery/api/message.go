package api

import "github.com/gin-gonic/gin"

func MessageRegister(r *gin.RouterGroup) {
	router := r.Group("messages")

	router.GET("")
	router.POST("")
	router.DELETE(":messageId")
}
