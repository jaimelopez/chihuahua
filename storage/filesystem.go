package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jaimelopez/chihuahua/executor"
)

// FileSystem todo
type FileSystem struct {
	file string
}

// NewFileSytemStorage todo
func NewFileSytemStorage(destination string) *FileSystem {
	return &FileSystem{
		file: destination,
	}
}

// GetLatest todo
func (fs *FileSystem) GetLatest() (*executor.Result, error) {
	result := &executor.Result{}

	content, err := ioutil.ReadFile(fs.file)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}

		return result, err
	}

	err = json.Unmarshal(content, result)

	return result, err
}

// Persist todo
func (fs *FileSystem) Persist(r *executor.Result) error {
	f, err := os.OpenFile(fs.file, os.O_CREATE|os.O_WRONLY, 0666)

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
