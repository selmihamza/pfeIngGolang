package Helpers

import (
	"driver/Models"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// UploadFile function
func UploadFile(r *http.Request, name string, id uint, typefile uint) Models.FileCar {
	var FileCar Models.FileCar
	file, handler, err := r.FormFile(name)
	if file != nil {
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)

		}
		defer file.Close()
		tempFile, err := ioutil.TempFile("public/vehicle/temp-images", name+"-*.png")
		nameFile := tempFile.Name()[strings.Index(tempFile.Name(), name+"-"):]
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		FileCar = Models.FileCar{Ref: nameFile, FileType: typefile, CreatedAt: time.Now(), UpdatedAt: time.Now(), DriverCarID: id, IsValid: false, Name: handler.Filename, Size: handler.Size}
	}
	return FileCar

}
