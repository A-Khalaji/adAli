package main

import (
    "adAli/internal/database"

    
	"github.com/gin-gonic/gin"
)

// @title           AdAli API
// @version         1.0
// @description     REST API for AdAli.
// @host            localhost:8000
// @BasePath        /
func main() {
    database.Connect()
    
	router := gin.Default()

	Routes(router)

	router.Run("localhost:8000")
}
