package entity

type VariableItemResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type VariableCreateRequest struct {
	Name          string `json:"name" validate:"required"`
	Value         string `json:"value" validate:"required"`
	Description   string `json:"description"`
	CloudRegionID uint   `json:"cloud_region_id"`
}

type VariableUpdateRequest struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Value       string `json:"value" validate:"required"`
	Description string `json:"description"`
}
