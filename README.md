# logz

Simple log wrapper to provide leveled, structured logging.

### Usage
```go
package main

import (
    "errors"
    "github.com/nuvuli/logz"
)

func main() {
    logger := logz.NewLogger(logz.Error) // only print errors

    logger.Info("key", "value", "key2", "value2") // not printed

    logger.InfoWithMessage("heyo") // "msg", "heyo" // not printed

    err := errors.New("something is wrong!")

    logger.Error(err, "key", "value") 
    // prints {"caller":"main.go:17","err":"something is wrong!","key":"value","level":"error","ts":"2019-09-30T13:33:26.253478Z"}
}
```