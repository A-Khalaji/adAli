package main

import (
	"adAli/internal/cache"
	"adAli/internal/database"
	"flag"
	"log"

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
	preloadCache := flag.Bool("load_cache", false, "load into redis")
	flag.Parse()

	database.Connect()
	database.ConnectRedis()


	if *preloadCache {
		if err := cache.Preload(); err != nil {
			log.Fatal(err)
		}
		log.Println("cache was loaded successfully")
	}

	router := gin.Default()

	Routes(router)

	router.Run("0.0.0.0:8000")
}
