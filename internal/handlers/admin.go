package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin godoc
//
// @Summary Admin panel
// @Tags Admin
// @Produce json
// @Success 200 {object} map[string]string
// @Router /admin [get]
func Admin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "admin panel",
	})
}