package cache

import (
	"adAli/internal/database"
	"adAli/internal/models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func SetZone(zone models.Zone) error {
	key := fmt.Sprintf("zone:%d", zone.ID)

	metadata, err := json.Marshal(zone.Metadata)
	if err != nil {
		return err
	}

	data := map[string]string{
		"id":         strconv.FormatUint(uint64(zone.ID), 10),
		"user_id":    strconv.FormatUint(uint64(zone.UserID), 10),
		"site_id":    strconv.FormatUint(uint64(zone.SiteID), 10),
		"name":       zone.Name,
		"zone_type":  string(zone.ZoneType),
		"identifier": zone.Identifier,
		"created_at": zone.CreatedAt.Format(time.RFC3339),
		"updated_at": zone.UpdatedAt.Format(time.RFC3339),
		"is_active":   strconv.FormatBool(zone.IsActive),
		"is_verified": strconv.FormatBool(zone.IsVerified),
		"metadata":   string(metadata),
	}

	return database.Redis.HSet(
		database.Ctx,
		key,
		data,
	).Err()
}

func DelZone(id uint) error {
	key := fmt.Sprintf("zone:%d", id)

	return database.Redis.Del(
		database.Ctx,
		key,
	).Err()
}

func UpZone(zone models.Zone) error {
	if err := DelZone(zone.ID); err != nil {
		return err
	}

	return SetZone(zone)
}
