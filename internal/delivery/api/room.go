package api

import "github.com/gin-gonic/gin"

func RoomRegister(r *gin.RouterGroup) {
	router := r.Group("rooms")

	router.GET("")
	router.POST("")
	router.DELETE(":roomId")
}
