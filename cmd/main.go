package main

import (
	"fmt"

	"github.com/dhawalhost/hooknetic/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("golang")
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
