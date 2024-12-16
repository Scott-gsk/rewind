package entity

type ConfigurationCreateRequest struct {
	Name            string `json:"name" validate:"required"`
	OperatingSystem string `json:"operating_system" validate:"required"`
	CollectorID     uint   `json:"collector_id" validate:"required"`
	CloudRegionID   uint   `json:"cloud_region_id" validate:"required"`
}

type ConfigurationUpdateRequest struct {
	ID              uint   `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	OperatingSystem string `json:"operating_system" validate:"required"`
	CollectorID     uint   `json:"collector_id" validate:"required"`
}

type ConfigurationItemResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	OperatingSystem string `json:"operating_system"`
	CollectorID     uint   `json:"collector_id"`
	CollectorName   string `json:"collector_name"`
	NodeCount       uint   `json:"node_count"`
}
