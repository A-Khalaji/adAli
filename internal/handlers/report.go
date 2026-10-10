package handlers

import (
	"adAli/internal/database"
	"adAli/internal/filter"
	"adAli/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReport godoc
//
//	@Summary		Get reports
//	@Description	Get all reports or filter and order them using query parameters.
//	@Description
//	@Description	Filtering operators:
//	@Description	:   (equals)
//	@Description	!:  (not equals)
//	@Description	>   (greater than)
//	@Description	>=  (greater than or equal)
//	@Description	<   (less than)
//	@Description	<=  (less than or equal)
//	@Description	~   (contains)
//	@Description
//	@Description	Examples:
//	@Description	/reports?filter=ad_id:15
//	@Description	/reports?filter=program_id:40
//	@Description	/reports?filter=date>=2026-09-01
//	@Description	/reports?filter=views>10
//	@Description	/reports?filter=clicks>5
//	@Description	/reports?order_by=views&sort=desc
//
//	@Tags			Report
//	@Produce		json
//
//	@Param			filter		query	[]string	false	"Filter expression. Can be repeated."
//	@Param			order_by	query	[]string	false	"Fields to order by. Can be repeated."
//	@Param			sort		query	[]string	false	"Sort direction for each order_by field (asc or desc). Can be repeated."
//
//	@Success		200		{array}		models.Report
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/reports [get]
func GetReport(c *gin.Context) {
	allowedFields := filter.AllowedFields(models.Report{})

	where, having, args, err := filter.ApplyClickHouse(c, allowedFields)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := filter.ApplyClickHouseOrder(c, allowedFields)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
		SELECT
			date,
			ad_id,
			program_id,
			zone_id,
			site_id,
			sum(reports.views) AS views,
			sum(reports.clicks) AS clicks,
			round(
				if(
					sum(reports.views) = 0,
					0,
					sum(reports.clicks) / sum(reports.views) * 100
				),
				2
			) AS click_rate
		FROM reports
	`

	if where != "" {
		query += "\n" + where
	}

	query += `
		GROUP BY
			date,
			ad_id,
			program_id,
			zone_id,
			site_id
	`

	if having != "" {
		query += "\n" + having
	}

	if order != "" {
		query += "\n" + order
	} else {
		query += "\nORDER BY date ASC, ad_id ASC"
	}

	reports := make([]models.Report, 0)
	
	if err := database.ClickHouseDB.Raw(query, args...).Scan(&reports).Error; err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reports)
}


