package handlers

import (
	"adAli/internal/database"
	"adAli/internal/filter"
	"adAli/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAds godoc
//
//	@Summary		Get ads
//	@Description	Get all ads or filter them using one or more `filter` query parameters.
//	@Description	Supported operators:
//	@Description	:   (equals)
//	@Description	!:  (not equals)
//	@Description	>   (greater than)
//	@Description	>=  (greater than or equal)
//	@Description	<   (less than)
//	@Description	<=  (less than or equal)
//	@Description	~   (contains, case-insensitive)
//	@Description
//	@Description	example: /ads?filter=is_active:true&filter=budget>=100000
//	@Tags			Ad
//	@Produce		json
//	@Param			filter	query	[]string	false	"Filter expression. Can be repeated."
//	@Success		200		{array}		models.Ad
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/ads [get]
func GetAds(c *gin.Context) {
	query := database.DB.Model(&models.Ad{}).Preload("User").Preload("Program")

	allowedFields := map[string]bool{
		"id":              true,
		"user_id":         true,
		"program_id":      true,
		"name":            true,
		"ad_type":         true,
		"destination_url": true,
		"created_at":      true,
		"is_active":       true,
		"is_verified":     true,
	}

	var err error

	query, err = filter.Apply(query, c, allowedFields)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var ads []models.Ad

	if err := query.Find(&ads).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ads)
}
