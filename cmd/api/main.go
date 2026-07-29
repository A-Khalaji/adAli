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
// @tag.name        Home
// @tag.name        Admin
// @tag.name        Publisher
// @tag.name        Advertiser
// @tag.name        Users
// @tag.name        Programs
// @tag.name        Ads
// @tag.name        Sites
// @tag.name        Zones
// @tag.name        Report
// @tag.name        Transaction
func main() {
	database.Connect()

	router := gin.Default()

	Routes(router)

	router.Run("localhost:8000")
}
