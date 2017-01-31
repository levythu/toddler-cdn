package simplememstorage

import (
    "storage"
    "sync"
)

// implement storage.StorageAPI
type SimpleMemoryStorage struct {
    m map[string][]byte
    lock sync.Mutex
}

func (_ *SimpleMemoryStorage)New(_ map[string]interface{}) storage.StorageAPI {
    return &SimpleMemoryStorage {
        m: map[string][]byte{},
    }
}

func (this *SimpleMemoryStorage)Put(key string, content []byte) error {
    this.lock.Lock()
    defer this.lock.Unlock()

    this.m[key]=content
    return nil
}

func (this *SimpleMemoryStorage)Get(key string) []byte {
    this.lock.Lock()
    defer this.lock.Unlock()

    return this.m[key]
}

func (this *SimpleMemoryStorage)Abandon() error {
    return nil
}
