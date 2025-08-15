package core

import (
	"mime/multipart"
)

type FileService interface {
	SaveFile(FileDTO FIleDTO, fileContent multipart.File) (FileResponse, error)
	FindFiles(FIleDTOGet FIleDTOGet) ([]FindFileResponse, error)
	Deletes(FIleDTOGet FIleDTOGet) ([]DeleteResponse, error)
}

type fileServiceImpl struct {
	repo FileRepository
}

func NewFileService(repo FileRepository) FileService {
	return &fileServiceImpl{
		repo: repo,
	}
}

func (f *fileServiceImpl) SaveFile(FileDTO FIleDTO, fileContent multipart.File) (FileResponse, error) {

	saveResult, err := f.repo.SaveFile(FileDTO, fileContent);
	if err != nil {
		return FileResponse{}, err
	}

 	return saveResult, nil;
};

func (f *fileServiceImpl) FindFiles(FIleDTOGet FIleDTOGet) ([]FindFileResponse, error) {
	getResult, err := f.repo.FindFiles(FIleDTOGet);

	if err != nil {
		return []FindFileResponse{}, err;
	}

	return getResult, nil;
}

func (f *fileServiceImpl) Deletes(FIleDTOGet FIleDTOGet) ([]DeleteResponse, error) {

	result, err := f.repo.Deletes(FIleDTOGet);
	if err != nil {
		return []DeleteResponse{}, nil;
	}

	return result, nil;
}

