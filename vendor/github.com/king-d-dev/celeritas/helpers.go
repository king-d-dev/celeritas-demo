package celeritas

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// func quitAppIfError(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func createDirIfNotExist(path string) error {
	const mode = 0755

	// if the path does not already exist, then go ahead and create it and return any error / non error(nil) return value you get
	// while trying to create the directory
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return os.Mkdir(path, mode)
	}
	// if the directory already exists, then do nothing and return nil
	return nil
}

func createFileIfNotExist(path string) error {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		file, err := os.Create(path)
		if err != nil {
			defer file.Close()
		}
		return err
	}
	return nil
}

func checkDotEnv(path string) error {
	return createFileIfNotExist(filepath.Join(path, ".env"))
}
