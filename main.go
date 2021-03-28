package main

import (
	"gingonic-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

}
