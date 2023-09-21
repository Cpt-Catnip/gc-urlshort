package urlshort

func buildMap(urls []mapping) map[string]string {
	pathToURLs := make(map[string]string)
	for _, pair := range urls {
		pathToURLs[pair.Path] = pair.URL
	}
	return pathToURLs
}
