package upload_api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
)

func UploadMultpleFile(response http.ResponseWriter, request *http.Request) {

	//Limit 10 MB
	request.ParseMultipartForm(10 * 1024 * 1024)
	files := request.MultipartForm.File["myfiles"]
	for _, file := range files {
		fmt.Println("File Info")
		fmt.Println("File Name:", file.Filename)
		fmt.Println("File Size:", file.Size)
		fmt.Println("File Type:", file.Header.Get("Content-Type"))
		fmt.Println("-------------------------")
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		f, _ := file.Open()

		tempFile, err := ioutil.TempFile("uploads", "*"+newFileName)
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		fileBytes, err2 := ioutil.ReadAll(f)
		if err2 != nil {
			fmt.Println(err2)
		}
		tempFile.Write(fileBytes)
	}
	fmt.Println("Done")
}

func HandleUpload(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20) // 10 MB
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		http.Error(response, "Bad request", http.StatusBadRequest)
		return
	}
	defer file.Close()
	log.Printf("Filename: %s\n", fileHeader.Filename)
	log.Printf("Size: %d\n", fileHeader.Size)
	log.Printf("Filename: %v\n", fileHeader.Header)

	// if fileHeader.Size >= 100000 {
	// 	log.Fatal("File to large")
	// }

	filename := path.Join("uploads", path.Base(fileHeader.Filename))
	dest, err := os.Create(filename)
	if err != nil {
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	// if _, err = io.Copy(dest, file); err != nil {
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	println("done")
}

func UploadFile(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("myfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("File Info")
	fmt.Println("File Name:", handler.Filename)
	fmt.Println("File Size:", handler.Size)
	fmt.Println("File Type:", handler.Header.Get("Content-Type"))

	extension := filepath.Ext(handler.Filename)
	fmt.Println(extension)
	newFileName := uuid.New().String() + extension
	fmt.Println(newFileName)
	tempFile, err2 := ioutil.TempFile("uploads", "*"+newFileName)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer tempFile.Close()
	fileBytes, err3 := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err3)
	}
	tempFile.Write(fileBytes)
	fmt.Println("Done")
}
