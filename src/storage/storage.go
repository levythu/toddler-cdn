package storage

type StorageAPI interface {
    New(context map[string]interface{}) StorageAPI

    Put(key string, content []byte) error
    Get(key string) []byte
}
