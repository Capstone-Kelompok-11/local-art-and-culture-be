package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IFilesService interface {
	CreateFiles(data *request.Files) (response.Files, error)
	GetAllFiles(nameFilter string) ([]response.Files, error)
	GetFiles(id string) (response.Files, error)
	UpdateFiles(id string, input request.Files) (response.Files, error)
	DeleteFiles(id string) (response.Files, error)
}

type FilesService struct {
	filesRepository repositories.IFilesRepository
}

func NewFilesService(repo repositories.IFilesRepository) *FilesService {
	return &FilesService{filesRepository: repo}
}

func (gu *FilesService) CreateFiles(data *request.Files) (response.Files, error) {
	if data.Image == "" {
		return response.Files{}, errors.ERR_IMAGE_IS_EMPTY
	}
	if data.SourceStr == "" || data.SourceId == 0 {
		return response.Files{}, errors.ERR_SOURCE_IS_EMPTY
	}

	res, err := gu.filesRepository.CreateFiles(data)
	if err != nil {
		return response.Files{}, errors.ERR_CREATE_FILES_DATABASE
	}
	return res, nil
}

func (gu *FilesService) GetAllFiles(nameFilter string) ([]response.Files, error) {
	res, err := gu.filesRepository.GetAllFiles(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (gu *FilesService) GetFiles(id string) (response.Files, error) {
	if id == "" {
		return response.Files{}, errors.ERR_GET_FILES_BAD_REQUEST_ID
	}
	res, err := gu.filesRepository.GetFiles(id)
	if err != nil {
		return response.Files{}, err
	}
	return res, nil
}

func (gu *FilesService) UpdateFiles(id string, input request.Files) (response.Files, error) {
	if id == "" {
		return response.Files{}, errors.ERR_GET_FILES_BAD_REQUEST_ID
	}
	res, err := gu.filesRepository.UpdateFiles(id, input)
	if err != nil {
		return response.Files{}, err
	}
	return res, nil
}

func (gu *FilesService) DeleteFiles(id string) (response.Files, error) {
	if id == "" {
		return response.Files{}, errors.ERR_GET_FILES_BAD_REQUEST_ID
	}
	res, err := gu.filesRepository.DeleteFiles(id)
	if err != nil {
		return response.Files{}, err
	}
	return res, nil
}
