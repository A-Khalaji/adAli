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

	users := router.Group("/users")
	{
		users.POST("", handlers.CreateUser)
		users.GET("", handlers.GetUsers)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	programs := router.Group("/programs")
	{
		programs.POST("", handlers.CreateProgram)
		programs.GET("", handlers.GetPrograms)
		programs.PUT("/:id", handlers.UpdateProgram)
		programs.DELETE("/:id", handlers.DeleteProgram)
	}

	ads := router.Group("/ads")
	{
		ads.POST("", handlers.CreateAd)
		ads.GET("", handlers.GetAds)
		ads.PUT("/:id", handlers.UpdateAd)
		ads.DELETE("/:id", handlers.DeleteAd)
	}

	sites := router.Group("/sites")
	{
		sites.POST("", handlers.CreateSite)
		sites.GET("", handlers.GetSite)
		sites.PUT("/:id", handlers.UpdateSite)
		sites.DELETE("/:id", handlers.DeleteSite)
	}

	zones := router.Group("/zones")
	{
		zones.POST("", handlers.CreateZone)
		zones.GET("", handlers.GetZone)
		zones.PUT("/:id", handlers.UpdateZone)
		zones.DELETE("/:id", handlers.DeleteZone)
	}

	reports := router.Group("/reports")
	{
		reports.GET("", handlers.GetReport)
	}

	transactions := router.Group("/transactions")
	{
		transactions.GET("", handlers.GetTransaction)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
