package diskstorage

import (
    "storage"
    "sync"
    . "logs"
    "crypto/md5"
    "fmt"
    "time"
)

const DEFAULT_ROOT_PATH="./diskstorage/"

// implement storage.StorageAPI
type DiskStorage struct {
    dstHost string

    storageFolder string
    currentSize int64
    cacheLiveTime int     // in seconds, -1 for forever, default is 1 hour

    lock sync.Mutex
    lockMap map[string]*sync.RWMutex
}

func (_ *DiskStorage)New(e map[string]interface{}) storage.StorageAPI {
    if err:=guaranteePathExist(DEFAULT_ROOT_PATH+"_"+e["dstHost"].(string)+"/"); err!=nil {
        L.Warn("Error when create dir", DEFAULT_ROOT_PATH+"_"+e["dstHost"].(string)+"/", err)
        return nil
    }
    var ret=&DiskStorage {
        dstHost: e["dstHost"].(string),
        storageFolder: DEFAULT_ROOT_PATH+"_"+e["dstHost"].(string)+"/",
        cacheLiveTime: 60*60,
        lockMap: make(map[string]*sync.RWMutex),
    }
    if v:=checkPathSize(ret.storageFolder); v<0 {
        L.Warn("Error when calculating the size of", ret.storageFolder)
        return nil
    } else {
        ret.currentSize=v
    }

    if v, ok:=e["TTL"].(float64); ok {
        ret.cacheLiveTime=int(v)
    }

    return ret
}

func (this *DiskStorage)putIfAbsent(key string, obj *sync.RWMutex) *sync.RWMutex {
    this.lock.Lock()
    defer this.lock.Unlock()
    if this.lockMap[key]==nil {
        this.lockMap[key]=obj
    }
    return this.lockMap[key]
}

func (this *DiskStorage)Put(key string, content []byte) error {
    key=fmt.Sprintf("%x", md5.Sum([]byte(key)))
    var lk=this.putIfAbsent(key, &sync.RWMutex{})
    lk.Lock()
    defer lk.Unlock()

    return writeFile(this.storageFolder+key+".dat", content, uint64(time.Now().Unix()))
}

func (this *DiskStorage)Get(key string) []byte {
    key=fmt.Sprintf("%x", md5.Sum([]byte(key)))
    var lk=this.putIfAbsent(key, &sync.RWMutex{})
    lk.Lock()
    defer lk.Unlock()

    if data, ts:=readFile(this.storageFolder+key+".dat"); data==nil {
        return nil
    } else if int64(ts)+int64(this.cacheLiveTime)<int64(time.Now().Unix()) {
        return nil
    } else {
        return data
    }
}

func (this *DiskStorage)Abandon() error {
    // TODO implement it!
    return nil
}
