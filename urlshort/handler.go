package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var handler http.Handler
		//if the url is found in our map..remap it
		if val, ok := pathsToUrls[r.URL.Path]; ok {
			handler = http.RedirectHandler(val, http.StatusMovedPermanently)
		} else {
			if fallback != nil {
				handler = fallback
			} else {
				handler = http.NotFoundHandler()
			}
		}

		handler.ServeHTTP(w, r)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYaml(yml []byte) ([]YamlFormatData, error) {
	var data []YamlFormatData
	err := yaml.Unmarshall(yml, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func buildMap(values []YamlFormatData) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, val := range values {
		if val.valid() {
			pathsToUrls[val.Path] = val.URL
		}
	}
	return pathsToUrls
}

type YamlFormatData struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
