package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jaimelopez/chihuahua/executor"
)

// FileSystemStorageExtension specifies the extension for stored files
const FileSystemStorageExtension string = ".bench"

// FileSystem struct representation
type FileSystem struct {
	destination string
	group       string
}

// NewFileSytemStorage driver
func NewFileSytemStorage(destination string, group string) (*FileSystem, error) {
	if fi, err := os.Stat(destination); os.IsNotExist(err) {
		return nil, err
	} else if !fi.IsDir() {
		return nil, errors.New("destination is not a valid directory")
	}

	return &FileSystem{
		destination: destination,
		group:       group,
	}, nil
}

// GetLatest stored results
func (fs *FileSystem) GetLatest() (*executor.Result, error) {
	result := &executor.Result{}

	content, err := ioutil.ReadFile(fs.filename())
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}

		return result, err
	}

	err = json.Unmarshal(content, result)

	return result, err
}

// Persist results
func (fs *FileSystem) Persist(r *executor.Result) error {
	f, err := os.OpenFile(fs.filename(), os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer f.Close()

	content, err := json.Marshal(r)
	if err != nil {
		return nil
	}

	_, err = f.Write(content)

	return err
}

func (fs *FileSystem) filename() string {
	return fmt.Sprint(
		fs.destination,
		string(os.PathSeparator),
		fs.group,
		FileSystemStorageExtension,
	)
}
