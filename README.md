# Golage log handling ![https://github.com/golage/log](https://godoc.org/github.com/golage/log?status.svg) ![https://github.com/golage/log](https://github.com/golage/log/workflows/Check/badge.svg) ![https://github.com/golage/log](https://codecov.io/gh/golage/log/branch/master/graph/badge.svg)
Simple and useful log handling package in Golang

## Installation
Get from Github:
```bash
go get github.com/golage/log
```

## How to use
Import into your code:
```go
import "github.com/golage/log"
```
Set output writer:
```go
log.SetOutput(w)
```
Set minimum log level (default: LevelInfo):
```go
log.SetLevel(log.LevelDebug)
```
Set log formatter (default: TextFormatter):
```go
log.SetFormatter(log.NewTextFormatter)
log.SetFormatter(log.NewJSONFormatter)
log.SetFormatter(log.NewYAMLFormatter)
```
Set constants data in all logs:
```go
log.SetConstant("key", "value")
```
Add data to log:
```go
log.With(err)
log.Value("key", "value")
```
Write log:
```go
log.Debug("message")
log.Info("message")
log.Warning("message")
log.Error("message")
log.Fatal("message")
```
For more see [example](examples/main.go)