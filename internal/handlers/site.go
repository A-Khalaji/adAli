package handlers

import (
	"adAli/internal/cache"
	"adAli/internal/database"
	"adAli/internal/filter"
	"adAli/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSite godoc
//
//	@Summary		Get sites
//	@Description	Get all sites or filter and order them using query parameters.
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
//	@Description	/sites?filter=is_active:true
//	@Description	/sites?filter=user_id:1
//	@Description	/sites?filter=name~example
//	@Description	/sites?order_by=created_at&sort=desc
//	@Description	/sites?order_by=name&order_by=created_at&sort=asc&sort=desc
//
//	@Tags			Sites
//	@Produce		json
//
//	@Param			filter		query	[]string	false	"Filter expression. Can be repeated."
//	@Param			order_by	query	[]string	false	"Fields to order by. Can be repeated."
//	@Param			sort		query	[]string	false	"Sort direction for each order_by field (asc or desc). Can be repeated."
//
//	@Success		200		{array}		models.Site
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/sites [get]
func GetSite(c *gin.Context) {
	query := database.DB.Model(&models.Site{}).Preload("User")

	allowedFields := filter.AllowedFields(models.Site{})

	var err error

	query, err = filter.Apply(query, c, allowedFields)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var sites []models.Site

	if err := query.Find(&sites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, sites)
}

// CreateSite godoc
//
//	@Summary		Create site
//	@Description	Create a new site.
//	@Tags			Sites
//	@Accept			json
//	@Produce		json
//
//	@Param			site	body		models.Site	true	"Site"
//
//	@Success		201		{object}	models.Site
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/sites [post]
func CreateSite(c *gin.Context) {
	var site models.Site

	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&site).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := cache.SetSite(site); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, site)
}

// UpdateSite godoc
//
//	@Summary		Update site
//	@Description	Update an existing site.
//	@Tags			Sites
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		int			true	"Site ID"
//	@Param			site	body		models.Site	true	"Updated site"
//
//	@Success		200		{object}	models.Site
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/sites/{id} [put]
func UpdateSite(c *gin.Context) {
    id := c.Param("id")
   
	var site models.Site
	if err := database.DB.First(&site, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "site not found",
		})
		return
	}
   
	var updates map[string]any
   
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
   
	if err := database.DB.Model(&site).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
   
	if err := database.DB.First(&site, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
   
	if err := cache.UpSite(site); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
   
	c.JSON(http.StatusOK, site)
}

// DeleteSite godoc
//
//	@Summary		Delete site
//	@Description	Delete a site by ID.
//	@Tags			Sites
//	@Produce		json
//
//	@Param			id	path		int	true	"Site ID"
//
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/sites/{id} [delete]
func DeleteSite(c *gin.Context) {
	id := c.Param("id")

	var site models.Site
	if err := database.DB.First(&site, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "site not found",
		})
		return
	}

	if err := database.DB.Delete(&site).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := cache.DelSite(site.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Site deleted successfully",
	})
}
