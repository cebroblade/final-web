package main

import (
	"os"
	router "web-final/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	loadEnv()
	server := gin.Default()
	db := connectDb()
	router := router.Router{Server: server, Db: db}
	router.Init()
	server.Run(os.Getenv("PORT"))
}
