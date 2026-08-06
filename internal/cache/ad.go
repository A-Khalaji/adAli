package cache

import (
	"adAli/internal/database"
	"adAli/internal/models"
	"fmt"
	"strconv"
	"time"
)

func SetAd(ad models.Ad) error {
	key := fmt.Sprintf("ad:%d", ad.ID)

	data := map[string]string{
		"id":          strconv.FormatUint(uint64(ad.ID), 10),
		"user_id":     strconv.FormatUint(uint64(ad.UserID), 10),
		"program_id":  strconv.FormatUint(uint64(ad.ProgramID), 10),
		"name":        ad.Name,
		"ad_type":     string(ad.AdType),
		"category":    ad.Metadata.Category,
		"keyword":     ad.Metadata.Keyword,
		"created_at":  ad.CreatedAt.Format(time.RFC3339),
		"updated_at":  ad.UpdatedAt.Format(time.RFC3339),
		"is_active":   strconv.FormatBool(ad.IsActive),
		"is_verified": strconv.FormatBool(ad.IsVerified),
	}

	if err := database.Redis.HSet(
		database.Ctx,
		key,
		data,
	).Err(); err != nil {
		return err
	}

	if err := database.Redis.SAdd(
		database.Ctx,
		"category:"+ad.Metadata.Category,
		ad.ID,
	).Err(); err != nil {
		return err
	}

	if err := database.Redis.SAdd(
		database.Ctx,
		"keyword:"+ad.Metadata.Keyword,
		ad.ID,
	).Err(); err != nil {
		return err
	}

	if err := database.Redis.SAdd(
		database.Ctx,
		"type:"+string(ad.AdType),
		ad.ID,
	).Err(); err != nil {
		return err
	}

	if ad.IsActive && ad.IsVerified {

		if err := database.Redis.SAdd(
			database.Ctx,
			"available_ads",
			ad.ID,
		).Err(); err != nil {
			return err
		}

	} else {

		if err := database.Redis.SRem(
			database.Ctx,
			"available_ads",
			ad.ID,
		).Err(); err != nil {
			return err
		}
	}

	return nil
}

func DelAd(ad models.Ad) error {
	if err := database.Redis.SRem(
		database.Ctx,
		"category:"+ad.Metadata.Category,
		ad.ID,
	).Err(); err != nil {
		return err
	}

	if err := database.Redis.SRem(
		database.Ctx,
		"keyword:"+ad.Metadata.Keyword,
		ad.ID,
	).Err(); err != nil {
		return err
	}

	if err := database.Redis.SRem(
		database.Ctx,
		"type:"+string(ad.AdType),
		ad.ID,
	).Err(); err != nil {
		return err
	}

	if err := database.Redis.SRem(
		database.Ctx,
		"available_ads",
		ad.ID,
	).Err(); err != nil {
		return err
	}

	key := fmt.Sprintf("ad:%d", ad.ID)

	return database.Redis.Del(
		database.Ctx,
		key,
	).Err()
}


func UpAd(oldAd, newAd models.Ad) error {

	if err := DelAd(oldAd); err != nil {
		return err
	}

	return SetAd(newAd)
}