package main

import (
    . "github.com/levythu/gurgling"
    . "logs"
    "os"
    "config"
)

func main() {
    L.Log("Launching from configuration file...")
    if err:=config.InitConfig(); err!=nil {
        L.Error("Fatal error encountered. Abort:", err)
        os.Exit(1)
    }
    if err:=InitAllServerConf(); err!=nil {
        L.Error("Fatal error encountered. Abort:", err)
        os.Exit(1)
    }
    var r=ARouter().Use(onRequest)
    L.Log("Launch server at port 80...")
    if err:=r.Launch(":80"); err!=nil {
        L.Error("Launch error, abort:", err)
        os.Exit(1)
    }
}

func onRequest(req Request, res Response) {

}
