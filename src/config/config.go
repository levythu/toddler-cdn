package config

import (
    "io/ioutil"
    "encoding/json"
    _ "fmt"
    . "logs"
)

var Conf map[string]interface{}

// Pay attention that filename is a relative path
func ReadFileToJSON(filename string) (map[string]interface{}, error) {
    var err error
    var res []byte

    res, err=ioutil.ReadFile(filename)
    if err!=nil {
        return nil, err
    }

    var ret map[string]interface{}
    err=json.Unmarshal(res, &ret)
    if err!=nil {
        return nil, err
    }

    return ret, nil
}

func InitConfig() error {
    var err error
    if Conf, err=ReadFileToJSON("./config.json"); err!=nil {
        L.Error("Error when loading config file:", err)
        return err
    }
    return nil
}
