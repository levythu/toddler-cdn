package main

import (
    // . "github.com/levythu/gurgling"
    . "logs"
    "config"
    "errors"
    "strconv"
)

type CDNServerConf struct {
    confName string

    srcHost string
    dstHost string
}

var confList []*CDNServerConf

func InitAllServerConf() error {
    var anonyCount=0
    confList=[]*CDNServerConf{}
    if list, ok:=config.Conf["servers"].([]interface{}); !ok {
        return errors.New("Incorrect tag 'servers'")
    } else {
        for _, oe:=range list {
            var e map[string]interface{}
            if e, ok=oe.(map[string]interface{}); !ok {
                continue
            }

            var newConf=&CDNServerConf{}
            var ok2 bool
            if newConf.confName, ok2=e["name"].(string); !ok2 {
                newConf.confName="Anonymous#"+strconv.Itoa(anonyCount)
                anonyCount++
            }
            if newConf.srcHost, ok2=e["srcHost"].(string); !ok2 {
                L.Warn(newConf.confName, "has incorrect tag 'srcHost'. Ignore it.")
                continue
            }
            if newConf.dstHost, ok2=e["dstHost"].(string); !ok2 {
                L.Warn(newConf.confName, "has incorrect tag 'dstHost'. Ignore it.")
                continue
            }
            confList=append(confList, newConf)
        }
    }

    return nil
}
