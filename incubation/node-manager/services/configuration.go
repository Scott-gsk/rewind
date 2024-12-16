package services

import (
	"fmt"
	"github.com/jinzhu/copier"
	"node-manager/entity"
	"node-manager/global"
	"node-manager/models"
)

type ConfigurationService struct {
}

func (s ConfigurationService) Create(req entity.ConfigurationCreateRequest) error {
	var configuration models.Configuration
	if err := copier.Copy(&configuration, req); err != nil {
		return err
	}
	if err := global.DBClient.Client.Create(&configuration).Error; err != nil {
		return err
	}
	return nil
}

func (s ConfigurationService) ListByCloudRegion(cloudRegionId uint) ([]entity.ConfigurationItemResponse, error) {
	var configurations []models.Configuration
	if err := global.DBClient.Client.
		Preload("Collector").
		Preload("Nodes").
		Where("cloud_region_id = ?", cloudRegionId).
		Find(&configurations).Error; err != nil {
		return nil, err
	}

	var configurationItemResponses []entity.ConfigurationItemResponse
	for _, config := range configurations {
		configurationItemResponses = append(configurationItemResponses, entity.ConfigurationItemResponse{
			ID:              config.ID,
			Name:            config.Name,
			OperatingSystem: config.OperatingSystem,
			CollectorID:     config.CollectorID,
			CollectorName:   config.Collector.Name,
			NodeCount:       uint(len(config.Nodes)),
		})
	}

	return configurationItemResponses, nil
}

func (s ConfigurationService) Update(req entity.ConfigurationUpdateRequest) error {
	var configuration models.Configuration
	if err := copier.Copy(&configuration, req); err != nil {
		return err
	}
	if err := global.DBClient.Client.Model(&configuration).
		Where("id = ?", req.ID).Updates(configuration).
		Error; err != nil {
		return err
	}
	return nil
}

func (s ConfigurationService) Delete(id uint) error {
	var config models.Configuration
	if err := global.DBClient.Client.Preload("Nodes").First(&config, id).Error; err != nil {
		return err
	}
	if len(config.Nodes) > 0 {
		return fmt.Errorf("不能删除已经应用的配置文件")
	}
	if err := global.DBClient.Client.Delete(&models.Configuration{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s ConfigurationService) DeleteMultiple(ids []uint) error {
	var configs []models.Configuration
	if err := global.DBClient.Client.Preload("Nodes").Where("id IN (?)", ids).Find(&configs).Error; err != nil {
		return err
	}
	for _, config := range configs {
		if len(config.Nodes) > 0 {
			return fmt.Errorf("cannot delete configuration that is applied to nodes")
		}
	}
	if err := global.DBClient.Client.Where("id IN (?)", ids).Delete(&models.Configuration{}).Error; err != nil {
		return err
	}
	return nil
}
