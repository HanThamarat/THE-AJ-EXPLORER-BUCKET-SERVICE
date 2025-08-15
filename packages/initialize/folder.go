package initialize

import "os"

func FolderInitialize() {
	err := os.MkdirAll("uploads", os.ModePerm);
	if err != nil {
		println(err);
		return;
	}	

	println("initialize the folder successfully!");
	return;
}