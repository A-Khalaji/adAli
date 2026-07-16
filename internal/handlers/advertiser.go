package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Advertiser godoc
//
// @Summary Advertiser panel
// @Tags Advertiser
// @Produce json
// @Success 200 {object} map[string]string
// @Router /advertiser [get]
func Advertiser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "advertiser panel",
	})
}