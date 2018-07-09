package storage

import "errors"

// Driver is used to queue asynchronous requests by a delegate middleware
type Driver string

// Storage drivers
const (
	FileSystemDriver    Driver = "file"
	ElasticSearchDriver Driver = "elastic"
)

// ErrInvalidStorageDriver when a unknown driver is specified
var ErrInvalidStorageDriver = errors.New("Storage driver missing or unknown")

// Storage todo
type Storage interface{}

// New todo
func New(driver Driver, destination string) (Storage, error) {
	switch driver {
	case FileSystemDriver:
		return NewFileSytemStorage(destination), nil
	case ElasticSearchDriver:
		return NewElasticSearchStorage(destination), nil
	default:
		return nil, ErrInvalidStorageDriver
	}
}
