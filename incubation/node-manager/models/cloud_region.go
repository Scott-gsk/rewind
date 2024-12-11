package models

import "gorm.io/gorm"

type CloudRegion struct {
	gorm.Model
	Name         string     `json:"name"`
	Introduction string     `json:"introduction"`
	Nodes        []Node     `gorm:"foreignKey:CloudRegionID"`
	Variables    []Variable `gorm:"foreignKey:CloudRegionID"`
}

func (CloudRegion) TableName() string {
	return "cloud_regions"
}
