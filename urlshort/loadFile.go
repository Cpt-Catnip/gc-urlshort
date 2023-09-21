package urlshort

import "os"

// LoadYAML takes a filepath as an argument and loads a yaml file from memory into a byte slice
func LoadYAML(f string) ([]byte, error) {
	return os.ReadFile(f)
}
