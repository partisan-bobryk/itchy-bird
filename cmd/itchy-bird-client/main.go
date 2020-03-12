package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	localBinRepo "github.com/VeprUA/itchy-bird/pkg/localbinaryrepository"
)

var DefaultFilePath string = "downloads"

func main() {
	binaryRepo, binaryRepoErr := localBinRepo.MakeLocalBinaryRepository(DefaultFilePath)

	if binaryRepoErr != nil {
		fmt.Println(binaryRepoErr)
		os.Exit(1)
	}

	// Get Arguments
	localVersionsArg := flag.Bool("ls", false, "Get list of local files")
	remoteVersionsArg := flag.Bool("ls-remote", false, "Get list of remote files")
	downloadBinary := flag.String("pull", "", "Pull down binary from remote\nUsage --pull <file-name>")

	// Execute command-line parsing
	flag.Parse()

	// Check local versions
	if *localVersionsArg == true {
		listOfHashedFiles := []localBinRepo.HashedFile{}

		fileList, fileListError := binaryRepo.GetListOfBinaries()
		if fileListError != nil {
			fmt.Println(fileListError)
			os.Exit(1)
		}

		for _, fileName := range fileList {
			hashedFile, hashedFileErr := binaryRepo.GetBinaryHash(fileName)

			if hashedFileErr != nil {
				// Inform the process but ommit from user payload
				fmt.Println(hashedFileErr)
				continue
			}

			// Add hashed file to an overall list
			listOfHashedFiles = append(listOfHashedFiles, hashedFile)
		}

		printVersions(listOfHashedFiles)

		// Exit gracefully
		os.Exit(0)
	}

	// Check remote versions
	if *remoteVersionsArg == true {
		versionList, err := getRemoteVersions()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		printVersions(versionList)

		// Exit gracefully
		os.Exit(0)
	}

	// Download file
	if downloadBinary != nil {
		downloadErr := downloadFile(*downloadBinary)

		if downloadErr != nil {
			fmt.Println(downloadErr)
			os.Exit(1)
		}
		// Exit gracefully
		os.Exit(0)
	}

}

func printVersions(versionList []localBinRepo.HashedFile) {
	for _, file := range versionList {
		fmt.Printf("%s %s\n", file.Name, file.Hash)
	}
}

func getRemoteVersions() ([]localBinRepo.HashedFile, error) {
	// Perform a fetch request
	// TODO Get a dedicated URL
	resp, respErr := http.Get("http://localhost:8080/versions")

	// Check for response errors
	if respErr != nil {
		return nil, respErr
	}

	// Close body because it is a stream
	defer resp.Body.Close()

	// Create a list of available versions
	var versionList []localBinRepo.HashedFile

	// Unmarshal JSON response to a struct
	jsonDecoderErr := json.NewDecoder(resp.Body).Decode(&versionList)

	// Check for any errors in decoding
	if jsonDecoderErr != nil {
		return nil, jsonDecoderErr
	}

	// Return list
	return versionList, nil
}

func downloadFile(filename string) error {
	jsonBodyString := []byte(`{}`)
	resp, respErr := http.Post("http://localhost:8080/download/"+filename, "application/json", bytes.NewBuffer(jsonBodyString))

	if respErr != nil {
		return respErr
	}

	defer resp.Body.Close()
	buffer, readAllErr := ioutil.ReadAll(resp.Body)

	if readAllErr != nil {
		return readAllErr
	}

	createFileErr := ioutil.WriteFile(DefaultFilePath+"/"+filename, buffer, 0644)

	if createFileErr != nil {
		return createFileErr
	}

	return nil
}
