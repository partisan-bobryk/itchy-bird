package main

import (
	"log"
	"os"
	"path/filepath"
)

const DEFAULT_FILE_PATH = "/downloads/*"

type BinaryRepository interface {
	GetListOfBinaries() (error, []string)
}

type LocalBinaryRepository struct {
	downloadLocation string
	isReady          bool
}

func (repo *LocalBinaryRepository) GetListOfBinaries() ([]string, error) {
	return nil, nil
}

func MakeLocalBinaryRepository(downloadLocation string) (*LocalBinaryRepository, error) {
	// Check for existance of local dir
	if _, err := os.Stat(downloadLocation); os.IsNotExist(err) {
		log.Printf("Could not find path %s, creating one...", downloadLocation)
		mkdirErr := os.Mkdir(downloadLocation, os.ModeDir)

		if mkdirErr != nil {
			log.Printf("Failed creating %s", downloadLocation)
			return nil, mkdirErr
		}
	}

	return &LocalBinaryRepository{
		downloadLocation: downloadLocation,
		isReady:          true,
	}, nil
}
