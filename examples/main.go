package main

import (
	"errors"

	"github.com/golage/log"
)

type data struct {
	Name string
}

func main() {
	log.SetLevel(log.LevelDebug)

	log.SetConstant("code_name", "example")

	log.Value("name", "john").Debug("debug message")

	log.With("john").Info("info message")

	log.With(data{Name: "john"}).Warning("warning message")

	log.With(errors.New("message")).Error("error message")

	log.Fatal("fatal error")
}
