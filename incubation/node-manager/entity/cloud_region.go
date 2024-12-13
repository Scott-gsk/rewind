package entity

type CloudRegionItemResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

type CloudRegionCreateRequest struct {
	Name         string `json:"name" validate:"required"`
	Introduction string `json:"introduction"`
}

type CloudRegionUpdateRequest struct {
	ID           uint   `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Introduction string `json:"introduction"`
}
