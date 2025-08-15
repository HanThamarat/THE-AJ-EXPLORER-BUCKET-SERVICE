package adapter

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	core "github.com/HanThamarat/GO-Bucket-Service/core/file"
	"github.com/google/uuid"
)

type fileSystemRepo struct {}

func NewFileRepository() core.FileRepository {
	return &fileSystemRepo{}
}

func (r *fileSystemRepo) SaveFile(FileDTO core.FIleDTO, data multipart.File) (core.FileResponse, error) {
	var fileData core.FIleDTO

	// Generate a unique filename using UUID
	uuid := uuid.New()
	fileExt := filepath.Ext(FileDTO.FileName);
	fileData.FileName = uuid.String() + "_" + strconv.Itoa(int(time.Now().Unix())) + fileExt;
	fileData.FilePath = FileDTO.FilePath

	err := os.MkdirAll("uploads" + fileData.FilePath, os.ModePerm)
	if err != nil {
		return core.FileResponse{}, err
	}

	filePath := filepath.Join("uploads" + fileData.FilePath, fileData.FileName)
	out, err := os.Create(filePath)
	if err != nil {
		return core.FileResponse{}, err
	}
	defer out.Close()

	// Copy the uploaded content to the file
	_, err = io.Copy(out, data)
	if err != nil {
		return core.FileResponse{}, err
	}

	// set the file create response.
	var response core.FileResponse;
	response.FileName = fileData.FileName;
	response.FilePath = "uploads" + fileData.FilePath + "/" + fileData.FileName;
	response.FileOriginalName = FileDTO.FileName;

	return response, nil;
};

func (r *fileSystemRepo) FindFiles(FileGet core.FIleDTOGet) ([]core.FindFileResponse, error) {
	var file core.FIleDTOGet;

	file.FilePath = FileGet.FilePath;
	file.FileName = FileGet.FileName;

	if len(file.FileName) == 0 {
		return []core.FindFileResponse{}, errors.New("Please enter file name you want to get");
	}

	var responses []core.FindFileResponse

	for _, value := range file.FileName {
		// Construct full path
		filePath := "uploads" + file.FilePath + "/" + value

		// Read file
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return []core.FindFileResponse{}, fmt.Errorf("error reading file %s: %w", value, err)
		}

		// Encode base64
		encodedString := base64.StdEncoding.EncodeToString(fileContent)

		response := core.FindFileResponse{
			FileName:        	value,
			FileOriginalName: 	value,
			FileBase94:      	encodedString,
			FilePath:        	file.FilePath,
		}
		responses = append(responses, response);
	}

	return responses, nil;
}

func (r *fileSystemRepo) Deletes(FIleDTOGet core.FIleDTOGet) ([]core.DeleteResponse, error) {
	var file core.FIleDTOGet;

	file.FilePath = FIleDTOGet.FilePath;
	file.FileName = FIleDTOGet.FileName;

	if len(file.FileName) == 0 {
		return []core.DeleteResponse{}, errors.New("Please enter file name you want to get");
	}

	var responses []core.DeleteResponse

	for _, fileName := range FIleDTOGet.FileName {
		
		filePath := filepath.Join("uploads", FIleDTOGet.FilePath, fileName)

		err := os.Remove(filePath)
		if err != nil {
			return nil, fmt.Errorf("error removing file %s: %w", fileName, err)
		}

		response := core.DeleteResponse{
			FileName: fileName,
			FilePath: FIleDTOGet.FilePath,
		}
		responses = append(responses, response);
	}

	return responses, nil;
}
