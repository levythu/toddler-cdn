package logs

import (
    "fmt"
    "sync"
)

type Logger interface {
    Log(c ...interface{})
    Warn(c ...interface{})
    Error(c ...interface{})
}

var L Logger=&consoleLogger{}

type consoleLogger struct {
    lock sync.Mutex
}

func (this *consoleLogger)Log(c ...interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    fmt.Print("[Log]  ")
    fmt.Println(c...)
}

func (this *consoleLogger)Warn(c ...interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    fmt.Print("[Warn] ")
    fmt.Println(c...)
}

func (this *consoleLogger)Error(c ...interface{}) {
    this.lock.Lock()
    defer this.lock.Unlock()

    fmt.Print("[Err]  ")
    fmt.Println(c...)
}
