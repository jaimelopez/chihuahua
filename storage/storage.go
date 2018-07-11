package storage

import (
	"errors"

	"github.com/jaimelopez/chihuahua/executor"
)

// Driver is used to queue asynchronous requests by a delegate middleware
type Driver string

// Storage drivers
const (
	FileSystemDriver    Driver = "file"
	ElasticSearchDriver Driver = "elastic"
)

// ErrInvalidStorageDriver when a unknown driver is specified
var ErrInvalidStorageDriver = errors.New("storage driver missing or unknown")

// Storage todo
type Storage interface {
	GetLatest() (*executor.Result, error)
	Persist(r *executor.Result) error
}

// New todo
func New(name string, driver string, destination string) (Storage, error) {
	switch Driver(driver) {
	case FileSystemDriver:
		return NewFileSytemStorage(destination), nil
	case ElasticSearchDriver:
		return NewElasticSearchStorage(destination, name), nil
	default:
		return nil, ErrInvalidStorageDriver
	}
}
