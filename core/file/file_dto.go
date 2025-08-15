package core

type FIleDTO struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}


type FIleDTOGet struct {
	FileName 	[]string 	`json:"file_name"`
	FilePath 	string 		`json:"file_path"`
}


type FileResponse struct {
	FileName 			string 			`json:"file_name"`
	FileOriginalName 	string 			`json:"file_original_name"`
	FilePath 			string 			`json:"file_path"`
}

type DeleteResponse struct {
	FileName 			string 			`json:"file_name"`
	FilePath 			string 			`json:"file_path"`
}

type FindFileResponse struct {
	FileName 			string 			`json:"file_name"`
	FileOriginalName 	string 			`json:"file_original_name"`
	FileBase94 			string   		`json:"file_base64"`
	FilePath 			string 			`json:"file_path"`
}