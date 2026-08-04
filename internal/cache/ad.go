package cache

import (
	"adAli/internal/database"
	"adAli/internal/models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func SetAd(ad models.Ad) error {
	key := fmt.Sprintf("ad:%d", ad.ID)

	metadata, err := json.Marshal(ad.Metadata)
	if err != nil {
		return err
	}

	data := map[string]string{
		"id":         strconv.FormatUint(uint64(ad.ID), 10),
		"user_id":    strconv.FormatUint(uint64(ad.UserID), 10),
		"program_id": strconv.FormatUint(uint64(ad.ProgramID), 10),
		"name":       ad.Name,
		"ad_type":    string(ad.AdType),
		"created_at": ad.CreatedAt.Format(time.RFC3339),
		"updated_at": ad.UpdatedAt.Format(time.RFC3339),
		"metadata":   string(metadata),
	}

	return database.Redis.HSet(
		database.Ctx,
		key,
		data,
	).Err()
}

func DelAd(id uint) error {
	key := fmt.Sprintf("ad:%d", id)

	return database.Redis.Del(
		database.Ctx,
		key,
	).Err()
}

func UpAd(ad models.Ad) error {
    if err := DelAd(ad.ID); err != nil{
        return err
    }

    return SetAd(ad)
}