package services

import (
	"github.com/jinzhu/copier"
	"node-manager/entity"
	"node-manager/global"
	"node-manager/models"
)

type CollectorService struct {
}

func (s CollectorService) Create(req entity.CollectorCreateRequest) error {
	var collector models.Collector
	if err := copier.Copy(&collector, req); err != nil {
		return err
	}
	if err := global.DBClient.Client.Create(&collector).Error; err != nil {
		return err
	}
	return nil
}

func (s CollectorService) List() ([]entity.CollectorItemResponse, error) {
	var collectorItemResponses []entity.CollectorItemResponse
	if err := global.DBClient.Client.
		Model(&models.Collector{}).
		Select("id, name, introduction, operating_system, details").
		Scan(&collectorItemResponses).Error; err != nil {
		return nil, err
	}
	return collectorItemResponses, nil
}

func (s CollectorService) Search(value string) ([]entity.CollectorItemResponse, error) {
	var collectorItemResponses []entity.CollectorItemResponse
	if err := global.DBClient.Client.
		Model(&models.Collector{}).
		Where("name LIKE ? OR introduction LIKE ?", "%"+value+"%", "%"+value+"%").
		Select("id, name, introduction, operating_system, details").
		Scan(&collectorItemResponses).Error; err != nil {
		return nil, err
	}
	return collectorItemResponses, nil
}
