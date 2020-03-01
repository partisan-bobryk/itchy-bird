package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiRouteHandlers struct {
	binaryRepository BinaryRepository
}

func (api *ApiRouteHandlers) GeVersionsHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")

	listOfHashedFiles := []HashedFile{}

	fileList, fileListError := api.binaryRepository.GetListOfBinaries()
	if fileListError != nil {
		// Inform logs
		log.Println(fileListError)
		// Hard Error! If we can't read the directory this should inform the client
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

	pyld := ResponsePayload{Status: http.StatusOK, Data: listOfHashedFiles}
	json.NewEncoder(res).Encode(pyld)
}

func (api *ApiRouteHandlers) DownloadHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	// TMP
	res.Header().Set("Content-Type", "application/json")
	pyld := ResponsePayload{Status: http.StatusOK, Data: "Looking good!"}

	json.NewEncoder(res).Encode(pyld)
}
