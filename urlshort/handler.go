package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

type mapping struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.String()
		if url, ok := pathsToUrls[path]; ok {
			// do redirect with new url
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			// use fallback
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml
	var urls []mapping
	if err := yaml.Unmarshal(yml, &urls); err != nil {
		return nil, err
	}

	// convert urls to map[string]string
	pathToURLs := buildMap(urls)

	return MapHandler(pathToURLs, fallback), nil
}

func JSONHandler(jsonBlob []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse json
	var urls []mapping
	if err := json.Unmarshal(jsonBlob, &urls); err != nil {
		return nil, err
	}

	// convert urls to map
	pathToUrls := buildMap(urls)

	return MapHandler(pathToUrls, fallback), nil
}

func FileHandler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	// load file into struct slice
	urls, err := loadFile(filename)
	if err != nil {
		return nil, err
	}

	// convert to map
	pathToUrls := buildMap(urls)

	// return handler
	return MapHandler(pathToUrls, fallback), nil
}
