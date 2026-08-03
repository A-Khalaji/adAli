package cache

import (
	"adAli/internal/database"
	"adAli/internal/models"
	"fmt"
	"strconv"
	"time"
)

func SetProgram(program models.Program) error {
	key := fmt.Sprintf("program:%d", program.ID)

	data := map[string]string{
		"id":          strconv.FormatUint(uint64(program.ID), 10),
		"user_id":     strconv.FormatUint(uint64(program.UserID), 10),
		"name":        program.Name,
		"description": program.Description,

		"budget":    strconv.FormatFloat(program.Budget, 'f', 2, 64),
		"bid_price": strconv.FormatFloat(program.BidPrice, 'f', 2, 64),

		"is_active":   strconv.FormatBool(program.IsActive),
		"is_verified": strconv.FormatBool(program.IsVerified),

		"created_at": program.CreatedAt.Format(time.RFC3339),
		"updated_at": program.UpdatedAt.Format(time.RFC3339),
	}

	if program.StartDate != nil {
		data["start_date"] = program.StartDate.Format(time.RFC3339)
	}

	if program.EndDate != nil {
		data["end_date"] = program.EndDate.Format(time.RFC3339)
	}

	return database.Redis.HSet(
		database.Ctx,
		key,
		data,
	).Err()
}
