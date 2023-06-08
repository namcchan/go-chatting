package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/namcchan/go-chatting/configs"
	"github.com/namcchan/go-chatting/database"
	"github.com/namcchan/go-chatting/internal/delivery/api"
	"github.com/namcchan/go-chatting/internal/delivery/ws"
	"github.com/namcchan/go-chatting/internal/middlewares"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	configs.LoadEnv()
	client, err := database.CreateMongoConnection()
	if err != nil {
		log.Fatal("Create mongodb connection error occurred")
		return
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			logrus.Panic(err)
			return
		}
	}(client, context.Background())

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RequestIDMiddleware())
	r.Use(middlewares.ErrorHandler())

	v1 := r.Group("api/v1")

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)

	api.AuthRegister(v1)
	api.RoomRegister(v1)
	api.AttachmentRegister(v1)

	r.NoRoute(middlewares.NotFoundHandler())

	_ = r.Run()
}
