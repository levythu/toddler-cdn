package simplememstorage

import (
    "storage"
)

// implement storage.StorageAPI
type SimpleMemoryStorage struct {
    m map[string][]byte
}

func (_ *SimpleMemoryStorage)New(_ map[string]interface{}) storage.StorageAPI {
    return &SimpleMemoryStorage {
        m: map[string][]byte{},
    }
}

func (this *SimpleMemoryStorage)Put(key string, content []byte) error {
    this.m[key]=content
    return nil
}

func (this *SimpleMemoryStorage)Get(key string) []byte {
    return this.m[key]
}
