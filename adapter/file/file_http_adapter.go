package adapter

import (
	core "github.com/HanThamarat/GO-Bucket-Service/core/file"
	"github.com/HanThamarat/GO-Bucket-Service/packages/response"
	"github.com/gofiber/fiber/v2"
)

type HttpFileAdapter struct {
	service core.FileService
}

func NewHttpFileHandler(service core.FileService) *HttpFileAdapter {
	return &HttpFileAdapter{
		service: service,
	}
}

// @Summary Upload a file
// @Description Upload a file and receive its metadata
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "File to upload"
// @Param file_path formData string true "path to save the files ex. /Han/png"
// @Success 201 {object} response.Response{data=core.FIleDTO} "File uploaded successfully"
// @Failure 400 {object} response.Response{} "Invalid request"
// @Router /upload [post]
func (h *HttpFileAdapter) SaveFile(c *fiber.Ctx) error {
	var fileBody core.FIleDTO	
	
	fileBody.FilePath = c.FormValue("file_path");

	if fileBody.FilePath == "" {
		return response.SendErrorHandler(c, fiber.StatusNotFound, "Save file Error.", "Please Enter the file path do u want to save this file.");
	}

	println(fileBody.FilePath);

	// Get the uploaded file
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return response.SendErrorHandler(c, fiber.StatusBadRequest, "File is required", err.Error())
	}

	fileBody.FileName = fileHeader.Filename

	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return response.SendErrorHandler(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer file.Close()

	result, err := h.service.SaveFile(fileBody, file);

	if err != nil {
		return response.SendErrorHandler(c, fiber.StatusInternalServerError, "Failed to save file", err.Error());
	}

	return response.SendResponseHandler(c, fiber.StatusCreated, "Creting the file to bucket successfully!", result);
}


// @Summary Getting files
// @Description Getting file from bucket.
// @Tags File
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param filebody body core.FIleDTOGet true "file body"
// @Success 201 {object} response.Response{data=core.FIleDTOGet} "File uploaded successfully"
// @Failure 400 {object} response.Response{} "Invalid request"
// @Router /findfiles [post]
func (h *HttpFileAdapter) Finds(c *fiber.Ctx) error {
	var fileBody core.FIleDTOGet;

	if err := c.BodyParser(&fileBody); err != nil {
		return response.SendErrorHandler(c, fiber.StatusBadRequest, "Invalid your request", err.Error());
	}

	result, err := h.service.FindFiles(fileBody);

	if err != nil {
		return response.SendErrorHandler(c, fiber.StatusInternalServerError, "Getting the files failed.", err.Error());
	}

	return response.SendResponseHandler(c, fiber.StatusOK, "Getting files from bucket successfully!", result)
}

// @Summary Deleting files
// @Description Deleting file in bucket.
// @Tags File
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param filebody body core.FIleDTOGet true "file body"
// @Success 201 {object} response.Response{data=core.FIleDTOGet} "File uploaded successfully"
// @Failure 400 {object} response.Response{} "Invalid request"
// @Router /deletefiles [post]
func (h  *HttpFileAdapter) DeleteFile(c *fiber.Ctx) error {
	var fileBody core.FIleDTOGet;

	if err := c.BodyParser(&fileBody); err != nil {
		return response.SendErrorHandler(c, fiber.StatusBadRequest, "Invalid your request", err.Error());
	}

	result, err := h.service.Deletes(fileBody);
	if err != nil {
		return response.SendErrorHandler(c, fiber.StatusInternalServerError, "Delete the files failed.", err.Error());
	}

	return response.SendResponseHandler(c, fiber.StatusOK, "Delete file in bucket successfully!", result);
}