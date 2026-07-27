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
//	@Description	Get all ads or filter and order them using query parameters.
//	@Description
//	@Description	Filtering operators:
//	@Description	:   (equals)
//	@Description	!:  (not equals)
//	@Description	>   (greater than)
//	@Description	>=  (greater than or equal)
//	@Description	<   (less than)
//	@Description	<=  (less than or equal)
//	@Description	~   (contains, case-insensitive)
//	@Description
//	@Description	Examples:
//	@Description	/ads?filter=is_active:true
//	@Description	/ads?filter=program_id:1
//	@Description	/ads?filter=bid_price>=1000
//	@Description	/ads?order_by=bid_price&sort=desc
//	@Description	/ads?order_by=created_at&order_by=bid_price&sort=desc&sort=asc
//
//	@Tags			Ad
//	@Produce		json
//
//	@Param			filter		query	[]string	false	"Filter expression. Can be repeated."
//	@Param			order_by	query	[]string	false	"Fields to order by. Can be repeated."
//	@Param			sort		query	[]string	false	"Sort direction for each order_by field (asc or desc). Can be repeated."
//
//	@Success		200		{array}		models.Ad
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/ads [get]
func GetAds(c *gin.Context) {
	query := database.DB.Model(&models.Ad{}).Preload("User").Preload("Program")

	allowedFields := filter.AllowedFields(models.Ad{})

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
