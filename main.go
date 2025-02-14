package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Cpt-Catnip/gc-urlshort/urlshort"
)

const DefaultYaml = "paths.yaml"

func main() {
	// read flags
	filePathPtr := flag.String("file", DefaultYaml, "filepath to url mappings")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	redirectHandler, err := urlshort.FileHandler(*filePathPtr, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", redirectHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
