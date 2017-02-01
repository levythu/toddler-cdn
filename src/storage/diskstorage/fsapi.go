package diskstorage

import (
    "os"
    "io/ioutil"
    "path/filepath"
    "encoding/binary"
)

func guaranteePathExist(path string) error {
    return os.MkdirAll(path, 0777)
}

// -1 indicates error
func checkPathSize(path string) int64 {
    var size int64=0
    var err=filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        size+=info.Size()
        return err
    })
    if err!=nil {
        return -1
    }
    return size
}

func writeFile(filename string, content []byte, timestamp uint64) error {
    var b=make([]byte, 8)
    binary.LittleEndian.PutUint64(b, timestamp)
    if file, err:=os.Create(filename); err!=nil {
        return err
    } else {
        file.Write(b)
        file.Write(content)
        file.Close()
        return nil
    }
}

// (nil, ) for error or noresult
func readFile(filename string) ([]byte, uint64) {
    if file, err:=os.Open(filename); err!=nil {
        return nil, 0
    } else {
        defer file.Close()
        if data, err:=ioutil.ReadAll(file); err!=nil {
            return nil, 0
        } else {
            if len(data)<8 {
                return nil, 0
            }
            var timestamp=binary.LittleEndian.Uint64(data[:8])
            return data[8:], timestamp
        }
    }
}
