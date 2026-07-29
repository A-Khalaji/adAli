package handlers

import (
	"adAli/internal/database"
	"adAli/internal/filter"
	"adAli/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetZone godoc
//
//	@Summary		Get zones
//	@Description	Get all zones or filter and order them using query parameters.
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
//	@Description	/zones?filter=is_active:true
//	@Description	/zones?filter=site_id:1
//	@Description	/zones?filter=name~banner
//	@Description	/zones?order_by=created_at&sort=desc
//	@Description	/zones?order_by=name&order_by=created_at&sort=asc&sort=desc
//
//	@Tags			Zones
//	@Produce		json
//
//	@Param			filter		query	[]string	false	"Filter expression. Can be repeated."
//	@Param			order_by	query	[]string	false	"Fields to order by. Can be repeated."
//	@Param			sort		query	[]string	false	"Sort direction for each order_by field (asc or desc). Can be repeated."
//
//	@Success		200		{array}		models.Zone
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/zones [get]
func GetZone(c *gin.Context) {
	query := database.DB.Model(&models.Zone{}).
		Preload("User").
		Preload("Site")

	allowedFields := filter.AllowedFields(models.Zone{})

	var err error

	query, err = filter.Apply(query, c, allowedFields)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var zones []models.Zone

	if err := query.Find(&zones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, zones)
}

// CreateZone godoc
//
//	@Summary		Create zone
//	@Description	Create a new zone.
//	@Tags			Zones
//	@Accept			json
//	@Produce		json
//
//	@Param			zone	body		models.Zone	true	"Zone"
//
//	@Success		201		{object}	models.Zone
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/zones [post]
func CreateZone(c *gin.Context) {
	var zone models.Zone

	if err := c.ShouldBindJSON(&zone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&zone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, zone)
}

// UpdateZone godoc
//
//	@Summary		Update zone
//	@Description	Update an existing zone.
//	@Tags			Zones
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		int			true	"Zone ID"
//	@Param			zone	body		models.Zone	true	"Updated zone"
//
//	@Success		200		{object}	models.Zone
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/zones/{id} [put]
func UpdateZone(c *gin.Context) {
	id := c.Param("id")

	var zone models.Zone
	if err := database.DB.First(&zone, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "zone not found",
		})
		return
	}

	var updatedZone models.Zone
	if err := c.ShouldBindJSON(&updatedZone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&zone).Updates(updatedZone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, zone)
}

// DeleteZone godoc
//
//	@Summary		Delete zone
//	@Description	Delete a zone by ID.
//	@Tags			Zones
//	@Produce		json
//
//	@Param			id	path		int	true	"Zone ID"
//
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/zones/{id} [delete]
func DeleteZone(c *gin.Context) {
	id := c.Param("id")

	var zone models.Zone
	if err := database.DB.First(&zone, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "zone not found",
		})
		return
	}

	if err := database.DB.Delete(&zone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Zone deleted successfully",
	})
}