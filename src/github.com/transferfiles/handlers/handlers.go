package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/transferfiles/constants"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Transfer file is up and running"))
}

func ReceiverHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(constants.MAX_MEMORY)
	fileNames := []string{}
	for key, _ := range r.MultipartForm.File {
		fileNames = append(fileNames, key)
	}

	successfullyUploadedFiles := []string{}
	for _, fileName := range fileNames {
		file, header, err := r.FormFile(fileName)
		if err != nil {
			fmt.Fprintln(w, err)
			w.Header().Set("Content-Type", "application/x-protobuf")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer file.Close()

		out, err := os.Create("/tmp/" + fileName)
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			w.Header().Set("Content-Type", "application/x-protobuf")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		successfullyUploadedFiles = append(successfullyUploadedFiles, header.Filename)
	}
	// TODO need to add the uploaded files as a part of meesage
	fmt.Fprintf(w, "Files uploaded successfully:")
}

func SendFileHandler(w http.ResponseWriter, r *http.Request) {
	request, err := uploadFile("http://localhost:3000/receiveFiles", "/Users/zeeshana/out.txt")
	if err != nil {
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to send files---"+err.Error())
		return
	}
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to send files---"+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Files sent successfully:")
}

func uploadFile(uri string, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(filepath.Base(path), filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
