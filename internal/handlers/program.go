package handlers

import (
	"adAli/internal/database"
	"adAli/internal/filter"
	"adAli/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPrograms godoc
//
//	@Summary		Get programs
//	@Description	Get all programs or filter them using one or more `filter` query parameters.
//	@Description	Supported operators:
//	@Description	:   (equals)
//	@Description	!:  (not equals)
//	@Description	>   (greater than)
//	@Description	>=  (greater than or equal)
//	@Description	<   (less than)
//	@Description	<=  (less than or equal)
//	@Description	~   (contains, case-insensitive)
//	@Description
//	@Description	example: /programs?filter=is_active:true&filter=budget>=100000
//	@Tags			Program
//	@Produce		json
//	@Param			filter	query	[]string	false	"Filter expression. Can be repeated."
//	@Success		200		{array}		models.Program
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/programs [get]
func GetPrograms(c *gin.Context) {
	query := database.DB.Model(&models.Program{}).Preload("User")

	allowedFields := map[string]bool{
		"id":          true,
		"user_id":     true,
		"name":        true,
		"budget":      true,
		"bid_price":   true,
		"start_date":  true,
		"end_date":    true,
		"created_at":  true,
		"is_active":   true,
		"is_verified": true,
	}

	var err error

	query, err = filter.Apply(query, c, allowedFields)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var programs []models.Program

	if err := query.Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, programs)
}
