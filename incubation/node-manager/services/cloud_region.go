package services

import (
	"github.com/jinzhu/copier"
	"node-manager/entity"
	"node-manager/global"
	"node-manager/models"
)

type CloudRegionService struct {
}

func (s CloudRegionService) List() ([]entity.CloudRegionItemResponse, error) {
	var cloudRegionResList []entity.CloudRegionItemResponse
	err := global.DBClient.Client.
		Model(&models.CloudRegion{}).
		Select("id, name, introduction").
		Scan(&cloudRegionResList).Error
	if err != nil {
		return nil, err
	}
	return cloudRegionResList, nil
}

func (s CloudRegionService) Create(req entity.CloudRegionCreateRequest) error {
	var cloudRegion models.CloudRegion
	err := copier.Copy(&cloudRegion, req)
	if err != nil {
		return err
	}
	err = global.DBClient.Client.Create(&cloudRegion).Error
	if err != nil {
		return err
	}
	return nil
}

func (s CloudRegionService) Update(req entity.CloudRegionUpdateRequest) error {
	var cloudRegion models.CloudRegion
	err := copier.Copy(&cloudRegion, req)
	if err != nil {
		return err
	}
	err = global.DBClient.Client.Model(&cloudRegion).Where("id = ?", req.ID).Updates(cloudRegion).Error
	if err != nil {
		return err
	}
	return nil
}
