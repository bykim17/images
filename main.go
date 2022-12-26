package main

import (
	"fmt"
	"net/http"
	"uploadfile/upload_api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/upload", upload_api.UploadFile).Methods("POST")
	router.HandleFunc("/api/uploadmultiple", upload_api.UploadMultpleFile).Methods("POST")
	router.HandleFunc("/api/uploadMysql", upload_api.HandleUpload).Methods("POST")
	err := http.ListenAndServe(":5050", router)
	if err != nil {
		fmt.Println(err)
	}
}
