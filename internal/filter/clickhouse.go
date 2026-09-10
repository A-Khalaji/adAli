package filter

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var clickHouseOperators = []string{
	"!:",
	">=",
	"<=",
	">",
	"<",
	"~",
	":",
}

var clickHouseOperatorSQL = map[string]string{
	":":  "=",
	"!:": "!=",
	">":  ">",
	">=": ">=",
	"<":  "<",
	"<=": "<=",
	"~":  "LIKE",
}

func ApplyClickHouse(c *gin.Context, allowedFields map[string]bool) (string, string, []any, error) {
	rawFilters := c.QueryArray("filter")

	var whereConditions []string
	var havingConditions []string
	var args []any

	for _, raw := range rawFilters {
		for _, value := range strings.Split(raw, ",") {
			value = strings.TrimSpace(value)

			if value == "" {
				continue
			}

			field, operator, filterValue, err := parseClickHouse(value)
			if err != nil {
				return "", "", nil, err
			}

			if !allowedFields[field] {
				return "", "", nil, fmt.Errorf("unknown field: %s", field)
			}

			sqlOperator := clickHouseOperatorSQL[operator]

			if operator == "~" {
				filterValue = "%" + filterValue + "%"
			}

			if field == "views" || field == "clicks" {
				havingConditions = append(
					havingConditions,
					fmt.Sprintf(
						"sum(reports.%s) %s ?",
						field,
						sqlOperator,
					),
				)
			} else {
				whereConditions = append(
					whereConditions,
					fmt.Sprintf("%s %s ?", field, sqlOperator),
				)
			}

			args = append(args, filterValue)
		}
	}

	var where string
	var having string

	if len(whereConditions) > 0 {
		where = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	if len(havingConditions) > 0 {
		having = "HAVING " + strings.Join(havingConditions, " AND ")
	}

	return where, having, args, nil
}

func ApplyClickHouseOrder(c *gin.Context, allowedFields map[string]bool) (string, error) {
	orderBy := c.QueryArray("order_by")
	sort := c.QueryArray("sort")

	if len(orderBy) == 0 {
		return "", nil
	}

	var orders []string

	for i, field := range orderBy {
		field = strings.TrimSpace(field)

		if !allowedFields[field] {
			return "", fmt.Errorf("unknown order field: %s", field)
		}

		direction := "asc"

		if i < len(sort) {
			direction = strings.ToLower(strings.TrimSpace(sort[i]))
		}

		if direction != "asc" && direction != "desc" {
			return "", fmt.Errorf("unknown order direction: %s", direction)
		}

		if field == "views" || field == "clicks" {
			orders = append(
				orders,
				fmt.Sprintf("sum(reports.%s) %s", field, direction),
			)
		} else {
			orders = append(
				orders,
				fmt.Sprintf("%s %s", field, direction),
			)
		}
	}

	return "ORDER BY " + strings.Join(orders, ", "), nil
}

func parseClickHouse(raw string) (string, string, string, error) {
	for _, operator := range clickHouseOperators {
		if strings.Contains(raw, operator) {
			parts := strings.SplitN(raw, operator, 2)

			if len(parts) != 2 {
				return "", "", "", fmt.Errorf("invalid filter: %s", raw)
			}

			field := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			if field == "" || value == "" {
				return "", "", "", fmt.Errorf("invalid filter: %s", raw)
			}

			return field, operator, value, nil
		}
	}

	return "", "", "", fmt.Errorf("invalid filter: %s", raw)
}