package main

import (
    // . "github.com/levythu/gurgling"
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
}
