package main

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

type record struct {
	Path string
	Url  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		dest, ok := pathsToUrls[path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, dest, http.StatusFound)
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
	urlMap, err := parseYAML(yml)
	if err != nil {
		return fallback.ServeHTTP, nil
	}

	return MapHandler(urlMap, fallback), nil
}

func parseYAML(yml []byte) (map[string]string, error) {
	var records []record
	err := yaml.Unmarshal(yml, &records)
	if err != nil {
		fmt.Println("failed to parse file " + err.Error())
		return nil, err
	}
	return convertToMap(records), nil
}

func convertToMap(records []record) map[string]string {
	urlMap := make(map[string]string, len(records))
	for _, rec := range records {
		urlMap[rec.Path] = rec.Url
	}
	return urlMap
}
