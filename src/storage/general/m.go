package general

import (
    . "storage"
    "storage/simplememstorage"
)

var nameTypeMap=map[string]StorageAPI{
    "simple-mem": (*simplememstorage.SimpleMemoryStorage)(nil),
}

func GetStorageAPI(storageName string, context map[string]interface{}) StorageAPI {
    if v, ok:=nameTypeMap[storageName]; ok {
        return v.New(context)
    } else {
        return nil
    }
}
