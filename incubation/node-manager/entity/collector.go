package entity

type CollectorItemResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Introduction    string `json:"description"`
	OperatingSystem string `json:"operating_system"`
	Details         string `json:"details"`
}

type CollectorCreateRequest struct {
	Name            string `json:"name" validate:"required"`
	Introduction    string `json:"description" validate:"required"`
	OperatingSystem string `json:"operating_system" validate:"required"`
	Details         string `json:"details" validate:"required"`
}
