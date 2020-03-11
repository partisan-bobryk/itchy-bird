package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiRouteHandlers struct {
	binaryRepository BinaryRepository
}

func (api *ApiRouteHandlers) GeVersionsHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	listOfHashedFiles := []HashedFile{}

	fileList, fileListError := api.binaryRepository.GetListOfBinaries()
	if fileListError != nil {
		log.Println(fileListError)
		// Hard Error! If we can't read the directory this should inform the client
		res.WriteHeader(http.StatusInternalServerError)
		pyld := ResponsePayload{http.StatusInternalServerError, fileListError}
		json.NewEncoder(res).Encode(pyld)
		return
	}

	for _, fileName := range fileList {
		hashedFile, hashedFileErr := api.binaryRepository.GetBinaryHash(fileName)

		if hashedFileErr != nil {
			// Inform the process but ommit from user payload
			log.Println(fileListError)
			continue
		}

		// Add hashed file to an overall list
		listOfHashedFiles = append(listOfHashedFiles, hashedFile)
	}

	res.WriteHeader(http.StatusOK)
	pyld := ResponsePayload{Status: http.StatusOK, Data: listOfHashedFiles}
	json.NewEncoder(res).Encode(pyld)
}

func (api *ApiRouteHandlers) DownloadHandler(res http.ResponseWriter, req *http.Request) {
	filename := mux.Vars(req)["fileName"]

	file, fileError := api.binaryRepository.GetFile(filename)
	defer file.Close()

	if fileError != nil {
		log.Println(fileError)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		pyld := ResponsePayload{Status: http.StatusNotFound, Data: fileError}
		json.NewEncoder(res).Encode(pyld)
		return
	}

	// Set header so that client can expect a binary
	res.Header().Set("Content-Type", "application/octet-stream")
	res.Header().Set("Content-Disposition", "attachment; filename="+filename)

	res.WriteHeader(http.StatusOK)
	io.Copy(res, file)
}
