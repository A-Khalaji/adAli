package cache

import (
	"adAli/internal/database"
	"adAli/internal/models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func SetSite(site models.Site) error {
	key := fmt.Sprintf("site:%d", site.ID)

	metadata, err := json.Marshal(site.Metadata)
	if err != nil {
	
		return err
	}

	data := map[string]string{
		"id":          strconv.FormatUint(uint64(site.ID), 10),
		"user_id":     strconv.FormatUint(uint64(site.UserID), 10),
		"name":        site.Name,
		"domain":      site.Domain,
		"identifier":  site.Identifier,
		"is_active":   strconv.FormatBool(site.IsActive),
		"is_verified": strconv.FormatBool(site.IsVerified),
		"created_at":  site.CreatedAt.Format(time.RFC3339),
		"updated_at":  site.UpdatedAt.Format(time.RFC3339),
		"metadata":    string(metadata),
	}

	return database.Redis.HSet(
		database.Ctx,
		key,
		data,
	).Err()
}

func DelSite(id uint) error {
	key := fmt.Sprintf("site:%d", id)

	return database.Redis.Del(
		database.Ctx,
		key,
	).Err()
}

func UpSite(site models.Site) error {
	if err := DelSite(site.ID); err != nil {
		return err
	}

	return SetSite(site)
}
