package main

import (
    . "github.com/levythu/gurgling"
    "net/http"
    "crypto/tls"
    "strings"
    "io/ioutil"
)

type IntermediateContent struct {
    Content []byte
    StatusCode int
    Header http.Header
}

var pipeClient=&http.Client{
    CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
            return http.ErrUseLastResponse
    },
}
var piptClientSSL=&http.Client{
    Transport: &http.Transport{
	       TLSClientConfig: &tls.Config{},
    },
    CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
            return http.ErrUseLastResponse
    },
}
func Pipe(req Request, targetURL string) (*IntermediateContent, error) {
    if strings.HasPrefix(strings.ToLower(targetURL), "https://") {
        return PipeX(req, targetURL, piptClientSSL)
    }
    return PipeX(req, targetURL, pipeClient)
}

func deepCopyHeader(src http.Header) http.Header {
    var srch=map[string][]string(src)
    var target=map[string][]string{}
    for k, v:=range srch {
        var tmp=make([]string, len(v))
        copy(tmp, v)
        target[k]=tmp
    }

    return target
}
func deepCopyHeaderIn(src http.Header, des http.Header) {
    var srch=map[string][]string(src)
    var target=map[string][]string(des)
    for k, v:=range srch {
        var tmp=make([]string, len(v))
        copy(tmp, v)
        target[k]=tmp
    }
}

func PipeX(req Request, targetURL string, client *http.Client) (*IntermediateContent, error) {
    var oReq=req.R()
    var proxyRequest, err=http.NewRequest(oReq.Method, targetURL, oReq.Body)
    if err!=nil {
        return nil, err
    }

    proxyRequest.Header=deepCopyHeader(oReq.Header)
    var proxyResponse, err2=client.Do(proxyRequest)
    if err2!=nil {
        return nil, err2
    }

    var theResult IntermediateContent

    theResult.Header=deepCopyHeader(proxyResponse.Header)
    theResult.StatusCode=proxyResponse.StatusCode
    if content, err:=ioutil.ReadAll(proxyResponse.Body); err!=nil {
        return nil, err
    } else {
        theResult.Content=content
    }
    proxyResponse.Body.Close()

    return &theResult, nil
}

func renderResponse(res Response, iData *IntermediateContent) {
    var oRes=res.R()
    deepCopyHeaderIn(iData.Header, oRes.Header())
    oRes.WriteHeader(iData.StatusCode)
    oRes.Write(iData.Content)
}
