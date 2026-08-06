package handlers

import (
	"adAli/internal/database"
	"adAli/internal/filter"
	"adAli/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
//
//	@Summary		Get users
//	@Description	Get all users or filter and order them using query parameters.
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
//	@Description	/users?filter=is_active:true
//	@Description	/users?filter=user_id:1
//	@Description	/users?filter=bid_price>=1000
//	@Description	/users?order_by=bid_price&sort=desc
//	@Description	/users?order_by=created_at&order_by=bid_price&sort=desc&sort=asc
//
//	@Tags			Users
//	@Produce		json
//
//	@Param			filter		query	[]string	false	"Filter expression. Can be repeated."
//	@Param			order_by	query	[]string	false	"Fields to order by. Can be repeated."
//	@Param			sort		query	[]string	false	"Sort direction for each order_by field (asc or desc). Can be repeated."
//
//	@Success		200		{array}		models.User
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
// @Router /users [get]
// @Router /users [post]
// @Router /users/{id} [put]
// @Router /users/{id} [delete]
func GetUsers(c *gin.Context) {
	query := database.DB.Model(&models.User{})

	allowedFields := filter.AllowedFields(models.User{})

	var err error

	query, err = filter.Apply(query, c, allowedFields)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var users []models.User

	if err := query.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
//
//	@Summary		Create user
//	@Description	Create a new user.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//
//	@Param			user	body		models.User	true	"User"
//
//	@Success		201		{object}	models.User
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/users [post]
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
//
//	@Summary		Update user
//	@Description	Update an existing user.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		int			true	"User ID"
//	@Param			user	body		models.User	true	"Updated user"
//
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/users/{id} [put]
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
   
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
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
   
	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete a user by ID.
//	@Tags			Users
//	@Produce		json
//
//	@Param			id	path		int	true	"User ID"
//
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
