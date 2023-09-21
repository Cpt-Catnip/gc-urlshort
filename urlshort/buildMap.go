package urlshort

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func loadFile(f string) ([]mapping, error) {
	// loadFile
	file, err := os.ReadFile(f)
	if err != nil {
		return []mapping{}, err
	}

	// unmarshal file
	var urls []mapping
	isYAML := strings.HasSuffix(f, ".yaml") || strings.HasSuffix(f, ".yml")
	isJSON := strings.HasSuffix(f, ".json")
	if isYAML {
		log.Print("loading yaml file")
		err := yaml.Unmarshal(file, &urls)
		if err != nil {
			return []mapping{}, err
		}
	} else if isJSON {
		log.Print("loading json file")
		err := json.Unmarshal(file, &urls)
		if err != nil {
			return []mapping{}, err
		}
	} else {
		log.Print("unrecognized file type")
		parts := strings.Split(f, ".")
		ext := parts[len(parts)-1]
		return []mapping{}, fmt.Errorf("unsupported filetype %q", ext)
	}

	return urls, nil
}

func buildMap(urls []mapping) map[string]string {
	pathToURLs := make(map[string]string)
	for _, pair := range urls {
		pathToURLs[pair.Path] = pair.URL
	}
	return pathToURLs
}
