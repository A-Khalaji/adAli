package filter

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Filter struct {
	Field    string
	Operator string
	Value    string
}

var operators = []string{
	"!:",
	">=",
	"<=",
	">",
	"<",
	"~",
	":",
}

var operatorSQL = map[string]string{
	":":  "=",
	"!:": "!=",
	">":  ">",
	">=": ">=",
	"<":  "<",
	"<=": "<=",
}

func parse(raw string) (Filter, error) {
	for _, op := range operators {
		if strings.Contains(raw, op) {
			parts := strings.SplitN(raw, op, 2)

			if len(parts) != 2 {
				return Filter{}, fmt.Errorf("invalid filter: %s", raw)
			}

			return Filter{
				Field:    strings.TrimSpace(parts[0]),
				Operator: op,
				Value:    strings.TrimSpace(parts[1]),
			}, nil
		}
	}
	return Filter{}, fmt.Errorf("invalid filter: %s", raw)
}

func Apply(query *gorm.DB, c *gin.Context, allowedFields map[string]bool) (*gorm.DB, error) {

	rawFilters := c.QueryArray("filter")

	var filters []string

	for _, raw := range rawFilters {

		for _, filter := range strings.Split(raw, ",") {

			filter = strings.TrimSpace(filter)

			if filter != "" {
				filters = append(filters, filter)
			}
		}
	}

	for _, raw := range filters {

		f, err := parse(raw)
		if err != nil {
			return nil, err
		}

		if !allowedFields[f.Field] {
			return nil, fmt.Errorf("unknown field: %s", f.Field)
		}

		if f.Operator == "~" {

			query = query.Where(
				fmt.Sprintf("%s ILIKE ?", f.Field),
				"%"+f.Value+"%",
			)

		} else {

			sqlOp, ok := operatorSQL[f.Operator]
			if !ok {
				return nil, fmt.Errorf("unknown operator: %s", f.Operator)
			}

			query = query.Where(
				fmt.Sprintf("%s %s ?", f.Field, sqlOp),
				f.Value,
			)
		}
	}

	return query, nil
}