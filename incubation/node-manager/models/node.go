package models

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	OperatingSystem string          `json:"operating_system"`
	IP              string          `json:"ip"`
	StatusID        uint            `json:"status_id"`
	Status          Status          `gorm:"foreignKey:StatusID"`
	CloudRegionID   uint            `json:"cloud_region_id"`
	CloudRegion     CloudRegion     `gorm:"foreignKey:CloudRegionID"`
	Configurations  []Configuration `gorm:"many2many:configuration_nodes;"` // many to many
}

type Status struct {
	gorm.Model
	CollectorStatusList []CollectorStatus `json:"collectors" gorm:"foreignKey:StatusID"`
	SidecarStatus       int               `json:"status"`
}

type CollectorStatus struct {
	gorm.Model
	StatusID        uint          `json:"status_id"`
	CollectorID     uint          `json:"collector_id"`
	Collector       Collector     `gorm:"foreignKey:CollectorID"`
	ConfigurationID uint          `json:"configuration_id"`
	Configuration   Configuration `gorm:"foreignKey:ConfigurationID"`
	CollectorStatus int           `json:"collector_status"`
}
