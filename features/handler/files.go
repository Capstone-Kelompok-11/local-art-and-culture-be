package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"
	"lokasani/helpers/upload"

	"github.com/labstack/echo"
)

type FilesHandler struct {
	filesService services.IFilesService
}

func NewFilesHandler(iFilesService services.IFilesService) *FilesHandler {
	return &FilesHandler{filesService: iFilesService}
}

func (fh *FilesHandler) CreateFiles(c echo.Context) error {
	var input request.Files
	c.Bind(&input)

	file, err := c.FormFile("image")
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	client := upload.ConfigCloud()
	imageUrl := upload.UploadFile(file, client)
	input.Url = imageUrl
	res, err := fh.filesService.CreateFiles(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (fh *FilesHandler) GetAllFiles(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	res, err := fh.filesService.GetAllFiles(nameFilter)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (fh *FilesHandler) GetFiles(c echo.Context) error {
	id := c.Param("id")
	res, err := fh.filesService.GetFiles(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (fh *FilesHandler) UpdateFiles(c echo.Context) error {
	id := c.Param("id")
	var input request.Files
	c.Bind(&input)

	res, err := fh.filesService.UpdateFiles(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (fh *FilesHandler) DeleteFiles(c echo.Context) error {
	id := c.Param("id")

	res, err := fh.filesService.DeleteFiles(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
