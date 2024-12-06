package main

import (
	"seanThakur/go-restapi/db"
	"seanThakur/go-restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
