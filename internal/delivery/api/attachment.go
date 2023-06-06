package api

import "github.com/gin-gonic/gin"

func AttachmentRegister(r *gin.RouterGroup) {
	router := r.Group("attachments")

	router.GET("")
	router.POST("")
	router.GET(":id")
	router.DELETE(":id")
}
