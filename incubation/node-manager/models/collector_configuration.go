package models

import "gorm.io/gorm"

type CollectorConfiguration struct {
	gorm.Model
	Name            string    `json:"name"`
	OperatingSystem string    `json:"operating_system"`
	NodeID          uint      `json:"node_id"`
	Node            Node      `gorm:"foreignKey:NodeID"`
	CollectorID     uint      `json:"collector_id"`
	Collector       Collector `gorm:"foreignKey:CollectorID"`
}

func (CollectorConfiguration) TableName() string {
	return "collector_configurations"
}
