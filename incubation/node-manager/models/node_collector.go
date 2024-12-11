package models

import "gorm.io/gorm"

type NodeCollector struct {
	gorm.Model
	NodeID                   uint   `json:"node_id"`
	CollectorConfigurationID uint   `json:"collector_configuration_id"`
	Status                   string `json:"status"`
}

func (NodeCollector) TableName() string {
	return "node_collectors"
}
