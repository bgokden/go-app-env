package main

import (
	"fmt"
	"net/http"

	goappenv "github.com/bgokden/go-app-env"
)

func RunWithEnv(goappenv goappenv.GoAppEnv) error {
	logger := goappenv.GetLogger()
	logger.Printf("I am running in env %v\n", goappenv.GetName())
	mux := goappenv.GetServeMux()
	mux.HandleFunc("/example0", Serve)
	return nil
}

const text = "server example from github\n"

func Serve(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, text)
}

func main() {
	err := RunWithEnv(goappenv.Base())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	select {}
}
