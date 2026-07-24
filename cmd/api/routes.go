package main

import (
    "adAli/internal/handlers"
    
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "adAli/cmd/api/docs"
)


func Routes(router *gin.Engine) {
	router.GET("/", handlers.Home)

	router.GET("/admin", handlers.Admin)
	router.GET("/publisher", handlers.Publisher)
	router.GET("/advertiser", handlers.Advertiser)
	
	router.GET("/programs", handlers.GetPrograms)

	router.GET("/ads", handlers.GetAds)

	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
