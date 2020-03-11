package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type HashedFile struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type BinaryRepository interface {
	GetListOfBinaries() ([]string, error)
	GetBinaryHash(fileName string) (HashedFile, error)
	GetFile(fileName string) (*os.File, error)
}

// Constructor
func MakeLocalBinaryRepository(downloadLocation string) (*LocalBinaryRepository, error) {
	// Check for existance of local dir
	if _, err := os.Stat(downloadLocation); os.IsNotExist(err) {
		log.Printf("Could not find path %s, creating one...", downloadLocation)
		mkdirErr := os.Mkdir(downloadLocation, 0777)

		if mkdirErr != nil {
			log.Print("Failed!\n")
			log.Println(mkdirErr)
			return nil, mkdirErr
		}

		log.Print("Success!")
	}

	return &LocalBinaryRepository{
		downloadLocation: downloadLocation,
		isReady:          true,
	}, nil
}

type LocalBinaryRepository struct {
	downloadLocation string
	isReady          bool
}

func (repo *LocalBinaryRepository) GetListOfBinaries() ([]string, error) {
	fileList := []string{}
	dirList, readDirErr := ioutil.ReadDir(repo.downloadLocation)

	if readDirErr != nil {
		return nil, readDirErr
	}

	for _, dirEntity := range dirList {
		// We should skip over any directories
		if dirEntity.IsDir() == true {
			continue
		}

		fileList = append(fileList, dirEntity.Name())
	}

	return fileList, nil
}

func (repo *LocalBinaryRepository) GetBinaryHash(fileName string) (HashedFile, error) {
	hashedFile := HashedFile{
		Name: fileName,
	}
	log.Printf("Attempting to load %s...", fileName)
	file, fileOpenErr := os.Open(repo.downloadLocation + "/" + fileName)
	defer file.Close()

	if fileOpenErr != nil {
		log.Print("Failed!\n")
		return hashedFile, fileOpenErr
	}

	log.Print("Success!\n")
	hash := sha256.New()
	if _, copyFileErr := io.Copy(hash, file); copyFileErr != nil {
		return hashedFile, copyFileErr
	}

	// Convert a byte array into string
	hashedFile.Hash = hex.EncodeToString(hash.Sum(nil))

	return hashedFile, nil
}

func (repo *LocalBinaryRepository) GetFile(fileName string) (*os.File, error) {
	// TODO Locate the file
	log.Printf("Attempting to load %s...", fileName)
	return os.Open(repo.downloadLocation + "/" + fileName)

}
