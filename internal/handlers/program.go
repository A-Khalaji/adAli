package handlers

import (
	"adAli/internal/cache"
	"adAli/internal/database"
	"adAli/internal/filter"
	"adAli/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPrograms godoc
//
//	@Summary		Get programs
//	@Description	Get all programs or filter and order them using query parameters.
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
//	@Description	/programs?filter=is_active:true
//	@Description	/programs?filter=user_id:1
//	@Description	/programs?filter=budget>=1000
//	@Description	/programs?order_by=budget&sort=desc
//	@Description	/programs?order_by=created_at&order_by=budget&sort=desc&sort=asc
//
//	@Tags			Programs
//	@Produce		json
//
//	@Param			filter		query	[]string	false	"Filter expression. Can be repeated."
//	@Param			order_by	query	[]string	false	"Fields to order by. Can be repeated."
//	@Param			sort		query	[]string	false	"Sort direction for each order_by field (asc or desc). Can be repeated."
//
//	@Success		200		{array}		models.Program
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/programs [get]
func GetPrograms(c *gin.Context) {
	query := database.DB.Model(&models.Program{}).Preload("User")

	allowedFields := filter.AllowedFields(models.Program{})

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

// CreateProgram godoc
//
//	@Summary		Create program
//	@Description	Create a new program.
//	@Tags			Programs
//	@Accept			json
//	@Produce		json
//
//	@Param			program	body		models.Program	true	"Program"
//
//	@Success		201		{object}	models.Program
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/programs [post]
func CreateProgram(c *gin.Context) {
	var program models.Program

	if err := c.ShouldBindJSON(&program); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := cache.SetProgram(program); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, program)
}

// UpdateProgram godoc
//
//	@Summary		Update program
//	@Description	Update an existing program.
//	@Tags			Programs
//	@Accept			json
//	@Produce		json
//
//	@Param			id			path		int				true	"Program ID"
//	@Param			program		body		models.Program	true	"Updated program"
//
//	@Success		200			{object}	models.Program
//	@Failure		400			{object}	map[string]string
//	@Failure		404			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/programs/{id} [put]
func UpdateProgram(c *gin.Context) {
	id := c.Param("id")

	var program models.Program
	if err := database.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "program not found",
		})
		return
	}

	var updatedProgram models.Program
	if err := c.ShouldBindJSON(&updatedProgram); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&program).Updates(updatedProgram).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := cache.UpProgram(program); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, program)
}

// DeleteProgram godoc
//
//	@Summary		Delete program
//	@Description	Delete a program by ID.
//	@Tags			Programs
//	@Produce		json
//
//	@Param			id	path		int	true	"Program ID"
//
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/programs/{id} [delete]
func DeleteProgram(c *gin.Context) {
	id := c.Param("id")

	var program models.Program
	if err := database.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "program not found",
		})
		return
	}

	if err := database.DB.Delete(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := cache.DelProgram(program.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Program deleted successfully",
	})
}
