package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Cpt-Catnip/gc-urlshort/urlshort"
)

const DefaultYaml = "paths.yaml"

func main() {
	// read flags
	filePathPtr := flag.String("file", DefaultYaml, "filepath to url mappings")

	// loadFile
	file, err := os.ReadFile(*filePathPtr)
	if err != nil {
		panic(err)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var redirectHandler http.HandlerFunc

	isYAML := strings.HasSuffix(*filePathPtr, ".yaml") || strings.HasSuffix(*filePathPtr, ".yml")
	isJSON := strings.HasSuffix(*filePathPtr, ".json")
	switch {
	case isYAML:
		// Build the YAMLHandler using the mapHandler as the
		// fallback
		redirectHandler, err = urlshort.YAMLHandler(file, mapHandler)
		if err != nil {
			panic(err)
		}
	case isJSON:
		redirectHandler, err = urlshort.JSONHandler(file, mapHandler)
		if err != nil {
			panic(err)
		}
	default:
		parts := strings.Split(*filePathPtr, ".")
		ext := parts[len(parts)-1]
		panic(fmt.Errorf("unsupported filetype %s", ext))
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
