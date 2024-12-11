package services

import (
	"github.com/jinzhu/copier"
	"node-manager/entity"
	"node-manager/global"
	"node-manager/models"
)

type VariableService struct {
}

func (service VariableService) ListByCloudRegion(cloudRegionId uint) ([]entity.VariableItemResponse, error) {
	var variableListRes []entity.VariableItemResponse
	err := global.DBClient.Client.
		Model(&models.Variable{}).
		Select("ID as id, Name as name, Value as value, Description as description").
		Where("cloud_region_id = ?", cloudRegionId).
		Scan(&variableListRes).Error
	if err != nil {
		return nil, err
	}
	return variableListRes, nil
}

func (service VariableService) SearchByCloudRegion(cloudRegionId uint, value string) ([]entity.VariableItemResponse, error) {
	var variableListRes []entity.VariableItemResponse
	err := global.DBClient.Client.
		Model(&models.Variable{}).
		Select("ID as id, Name as name, Value as value, Description as description").
		Where("cloud_region_id = ? AND (Name LIKE ? OR Description LIKE ?)", cloudRegionId, "%"+value+"%", "%"+value+"%").
		Scan(&variableListRes).Error
	if err != nil {
		return nil, err
	}
	return variableListRes, nil
}

func (service VariableService) CreateVariable(req entity.VariableCreateRequest) error {
	var variable models.Variable
	err := copier.Copy(&variable, req)
	if err != nil {
		return err
	}
	err = global.DBClient.Client.Create(&variable).Error
	if err != nil {
		return err
	}
	return nil
}

func (service VariableService) UpdateVariable(request entity.VariableUpdateRequest) error {
	var variable models.Variable
	err := copier.Copy(&variable, request)
	if err != nil {
		return err
	}
	err = global.DBClient.Client.Model(&variable).Where("id = ?", request.ID).Updates(variable).Error
	if err != nil {
		return err
	}
	return nil
}

func (service VariableService) DeleteSingleVariable(u uint) error {
	err := global.DBClient.Client.Delete(&models.Variable{}, u).Error
	if err != nil {
		return err
	}
	return nil
}

func (service VariableService) DeleteMultipleVariables(ids []uint) error {
	err := global.DBClient.Client.Where("id IN (?)", ids).Delete(&models.Variable{}).Error
	if err != nil {
		return err
	}
	return nil
}
