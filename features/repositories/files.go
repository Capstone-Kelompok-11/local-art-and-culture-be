package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IFilesRepository interface {
	CreateFiles(data *request.Files) (response.Files, error)
	GetAllFiles(nameFilter string) ([]response.Files, error)
	GetFiles(id string) (response.Files, error)
	UpdateFiles(id string, input request.Files) (response.Files, error)
	DeleteFiles(id string) (response.Files, error)
}

type filesRepository struct {
	db *gorm.DB
}

func NewFilesRepository(db *gorm.DB) *filesRepository {
	return &filesRepository{db}
}

func (rr *filesRepository) CreateFiles(data *request.Files) (response.Files, error) {
	dataFiles := domain.ConvertFromFilesReqToModel(*data, data.Image)
	err := rr.db.Create(&dataFiles).Error
	if err != nil {
		return response.Files{}, err
	}
	
	return *domain.ConvertFromModelToFilesRes(*dataFiles), nil
}

func (rr *filesRepository) GetAllFiles(nameFilter string) ([]response.Files, error) {
	var allFiles []models.Files
	var resAllFiles []response.Files

	query := rr.db.Model(&models.Files{})
	if nameFilter != "" {
		query = query.Where("Files LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allFiles).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allFiles); i++ {
		FilesVm := domain.ConvertFromModelToFilesRes(allFiles[i])
		resAllFiles = append(resAllFiles, *FilesVm)
	}
	return resAllFiles, nil
}

func (rr *filesRepository) GetFiles(id string) (response.Files, error) {
	var filesData models.Files
	err := rr.db.First(&filesData, "id = ?", id).Error

	if err != nil {
		return response.Files{}, err
	}

	return *domain.ConvertFromModelToFilesRes(filesData), nil
}

func (rr *filesRepository) UpdateFiles(id string, input request.Files) (response.Files, error) {
	filesData := models.Files{}
	err := rr.db.First(&filesData, "id = ?", id).Error

	if err != nil {
		return response.Files{}, errors.ERR_GET_FILES_BAD_REQUEST_ID
	}

	if input.Image != "" {
		filesData.Image = input.Image
		
	if filesData.SourceId != 0 {
		filesData.SourceId = input.SourceId
	}
	if filesData.SourceStr != "" {
		filesData.SourceStr = input.SourceStr
	}

	if err = rr.db.Save(&filesData).Error; err != nil {
		return response.Files{}, errors.ERR_UPDATE_DATA
	}
}
	return *domain.ConvertFromModelToFilesRes(filesData), nil
}

func (rr *filesRepository) DeleteFiles(id string) (response.Files, error) {
	filesData := models.Files{}
	res := response.Files{}
	find := rr.db.First(&filesData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToFilesRes(filesData)
	}
	err := rr.db.Delete(&filesData, "id = ?", id).Error
	if err != nil {
		return response.Files{}, err
	}
	return res, nil
}