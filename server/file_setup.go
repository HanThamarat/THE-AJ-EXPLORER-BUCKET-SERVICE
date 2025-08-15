package server

import (
	adapter "github.com/HanThamarat/GO-Bucket-Service/adapter/file"
	core "github.com/HanThamarat/GO-Bucket-Service/core/file"
	"github.com/gofiber/fiber/v2"
)

func (c *fiberServer) InitializeFileService(api fiber.Router) {
	fileRepo := adapter.NewFileRepository();
	fileService := core.NewFileService(fileRepo);
	FileHandler := adapter.NewHttpFileHandler(fileService);

	api.Post("/upload", FileHandler.SaveFile);
	api.Post("/findfiles", FileHandler.Finds);
	api.Post("/deletefiles", FileHandler.DeleteFile);
}