package models

import "gorm.io/gorm"

type Variable struct {
	gorm.Model
	Name          string      `json:"name"`
	Value         string      `json:"value"`
	Description   string      `json:"description"`
	CloudRegionID uint        `json:"cloud_region_id"`
	CloudRegion   CloudRegion `gorm:"foreignKey:CloudRegionID"`
}
