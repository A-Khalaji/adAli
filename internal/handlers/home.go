package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home godoc
//
// @Summary Home
// @Tags Home
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to AdAli",
	})
}