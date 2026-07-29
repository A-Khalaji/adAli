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
//	@Tags			Ads
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
	query := database.DB.Model(&models.Ad{}).
		Preload("User").
		Preload("Program")

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

// CreateAd godoc
//
//	@Summary		Create ad
//	@Description	Create a new ad.
//	@Tags			Ads
//	@Accept			json
//	@Produce		json
//
//	@Param			ad	body		models.Ad	true	"Ad"
//
//	@Success		201		{object}	models.Ad
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/ads [post]
func CreateAd(c *gin.Context) {
	var ad models.Ad

	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&ad).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ad)
}

// UpdateAd godoc
//
//	@Summary		Update ad
//	@Description	Update an existing ad.
//	@Tags			Ads
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		int			true	"Ad ID"
//	@Param			ad	body		models.Ad	true	"Updated ad"
//
//	@Success		200		{object}	models.Ad
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/ads/{id} [put]
func UpdateAd(c *gin.Context) {
	id := c.Param("id")

	var ad models.Ad
	if err := database.DB.First(&ad, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ad not found",
		})
		return
	}

	var updatedAd models.Ad
	if err := c.ShouldBindJSON(&updatedAd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&ad).Updates(updatedAd).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ad)
}

// DeleteAd godoc
//
//	@Summary		Delete ad
//	@Description	Delete an ad by ID.
//	@Tags			Ads
//	@Produce		json
//
//	@Param			id	path		int	true	"Ad ID"
//
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/ads/{id} [delete]
func DeleteAd(c *gin.Context) {
	id := c.Param("id")

	var ad models.Ad
	if err := database.DB.First(&ad, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ad not found",
		})
		return
	}

	if err := database.DB.Delete(&ad).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ad deleted successfully",
	})
}