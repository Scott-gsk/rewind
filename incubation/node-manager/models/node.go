package models

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	Ip              string                   `json:"ip"`
	OperatingSystem string                   `json:"operating_system"`
	SidecarStatus   string                   `json:"sidecar_status"`
	CloudRegionID   uint                     `json:"cloud_region_id"`
	CloudRegion     CloudRegion              `gorm:"foreignKey:CloudRegionID"`
	Configurations  []CollectorConfiguration `gorm:"many2many:node_collector_configurations;"`
}

func (Node) TableName() string {
	return "nodes"
}
