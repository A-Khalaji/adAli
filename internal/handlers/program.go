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
