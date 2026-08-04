package cache

import (
	"adAli/internal/database"
	"adAli/internal/models"
)

func Preload() error {
	var programs []models.Program
	if err := database.DB.Find(&programs).Error; err != nil {
		return err
	}

	for _, program := range programs {
		if err := SetProgram(program); err != nil {
			return err
		}
	}

	var ads []models.Ad
	if err := database.DB.Find(&ads).Error; err != nil {
		return err
	}

	for _, ad := range ads {
		if err := SetAd(ad); err != nil {
			return err
		}
	}

	return nil
}
