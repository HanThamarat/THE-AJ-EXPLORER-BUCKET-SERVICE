package core

import "mime/multipart"

type FileRepository interface {
	SaveFile(FIleDTO FIleDTO, fileContent multipart.File) (FileResponse, error)
	FindFiles(FIleDTOGet FIleDTOGet) ([]FindFileResponse, error)
	Deletes(FIleDTOGet FIleDTOGet) ([]DeleteResponse, error)
}
