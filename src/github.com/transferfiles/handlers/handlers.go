package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Transfer file is up and running"))
}

var multipartByReader = &multipart.Form{
	Value: make(map[string][]string),
	File:  make(map[string][]*multipart.FileHeader),
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fileNames := []string{}
	for key, _ := range r.MultipartForm.File {
		fileNames = append(fileNames, key)
	}

	successfullyUploadedFiles := []string{}
	for _, fileName := range fileNames {
		file, header, err := r.FormFile(fileName)

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		defer file.Close()

		out, err := os.Create("/tmp/" + fileName)
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		successfullyUploadedFiles = append(successfullyUploadedFiles, header.Filename)
	}

	fmt.Fprintf(w, "Files uploaded successfully:")
}
