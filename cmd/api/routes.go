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
	router.GET("/zone", handlers.GetZone)
	router.GET("/site", handlers.GetSite)

	users := router.Group("/users")
	{
		users.GET("", handlers.GetUsers)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	router.GET("/report", handlers.GetReport)
	router.GET("/transaction", handlers.GetTransaction)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
