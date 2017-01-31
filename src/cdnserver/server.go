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

    go func() {
        var r=ARouter().Use(onRequest)
        L.Log("Launch server at port 80...")
        if err:=r.Launch(":80"); err!=nil {
            L.Error("Launch error, abort:", err)
            os.Exit(1)
        }
    } ()

    select{}
}

func onRequest(req Request, res Response) {
    if conf, ok:=confMap[req.R().Host]; ok {
        var desURL=*(req.R().URL)
        desURL.Scheme="http"
        desURL.Host=conf.dstHost
        if iData, err:=Pipe(req, desURL.String()); err!=nil {
            res.Status("Failed to get:"+err.Error(), 500)
        } else {
            renderResponse(res, iData)
        }
    } else {
        res.Status("Invalid hostname", 400)
    }
}

func onRequestHTTPS(req Request, res Response) {
    if conf, ok:=confMap[req.R().Host]; ok {
        var desURL=*(req.R().URL)
        desURL.Scheme="https"
        desURL.Host=conf.dstHost
        if iData, err:=Pipe(req, desURL.String()); err!=nil {
            res.Status("Failed to get:"+err.Error(), 500)
        } else {
            renderResponse(res, iData)
        }
    } else {
        res.Status("Invalid hostname", 400)
    }
}
