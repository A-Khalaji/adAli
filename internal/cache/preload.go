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
	
	var sites []models.Site
	if err := database.DB.Find(&sites).Error; err != nil {
		return err
	}

	for _, site := range sites {
		if err := SetSite(site); err != nil {
			return err
		}
	}
	
	var zones []models.Zone
	if err := database.DB.Find(&zones).Error; err != nil {
		return err
	}

	for _, zone := range zones {
		if err := SetZone(zone); err != nil {
			return err
		}
	}

	return nil
}
