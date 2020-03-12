package main

import (
	"os"

	localBinRepo "github.com/VeprUA/itchy-bird/pkg/localbinaryrepository"
)

type BinaryRepository interface {
	GetListOfBinaries() ([]string, error)
	GetBinaryHash(fileName string) (localBinRepo.HashedFile, error)
	GetFile(fileName string) (*os.File, error)
}
