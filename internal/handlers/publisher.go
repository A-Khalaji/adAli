package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Publisher godoc
//
// @Summary Publisher panel
// @Tags Publisher
// @Produce json
// @Success 200 {object} map[string]string
// @Router /publisher [get]
func Publisher(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "publisher panel",
	})
}