package main

import (
	"encoding/json"
	"net/http"
)

func GeVersionsHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	pyld := ResponsePayload{Status: http.StatusOK, Data: "Looking good!"}
	json.NewEncoder(res).Encode(pyld)
}

func DownloadHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	// TMP
	res.Header().Set("Content-Type", "application/json")
	pyld := ResponsePayload{Status: http.StatusOK, Data: "Looking good!"}

	json.NewEncoder(res).Encode(pyld)
}
