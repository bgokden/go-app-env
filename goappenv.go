package goappenv

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/bgokden/go-app-env/cache"
	"github.com/bgokden/go-app-env/kv"
)

type GoAppEnv interface {
	GetName() string
	GetEnvironmentVariable(string) string
	GetServeMux() *http.ServeMux
	GetCache() cache.Cache
	GetDB() *sql.DB
	GetKV() kv.KV
	GetLogger() *log.Logger
}

type GoApp interface {
	GetName() string
	GetDependencies() []string
	RunWithEnv(GoAppEnv) error
}

func Base() GoAppEnv {
	e := new(BaseGoAppEnv)
	go func() {
		err := http.ListenAndServe(":8080", e.GetServeMux())
		log.Fatal(err)
	}()
	e.Logger = log.New(os.Stdout, "GoAppEnv: ", log.Lshortfile)
	return e
}

type BaseGoAppEnv struct {
	Logger *log.Logger
}

func (e *BaseGoAppEnv) GetName() string {
	return "BaseGoAppEnv"
}

func (e *BaseGoAppEnv) GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

func (e *BaseGoAppEnv) GetServeMux() *http.ServeMux {
	return http.DefaultServeMux
}

func (e *BaseGoAppEnv) GetCache() cache.Cache {
	return nil
}

func (e *BaseGoAppEnv) GetDB() *sql.DB {
	return nil
}

func (e *BaseGoAppEnv) GetKV() kv.KV {
	return nil
}

func (e *BaseGoAppEnv) GetLogger() *log.Logger {
	return e.Logger
}

/*
example:

package main

import (
	"fmt"

	goappenv "github.com/bgokden/go-app-env"
)

func RunWithEnv(goappenv goappenv.GoAppEnv) error {
  fmt.Printf("I am running in env %v\n",goappenv.GetName())
  return nil
}


func main() {
  err := RunWithEnv(goappenv.Base())
}
*/
