package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Entry represents one of the entries in our json
type Entry struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	}
}

// HTMLHandler this function serves up the appropriate json object inside
// a templated html page
func HTMLHandler(json map[string]*Entry, key string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content Type", "text/html")
		tmpl := template.Must(template.ParseFiles("bookpage.html"))
		fmt.Println(r.URL.Path)
		if r.URL.Path[1:] == "" {
			r.URL.Path = key
		}
		if val, ok := json[r.URL.Path[1:]]; ok {
			tmpl.Execute(w, val)
		} else {
			handler := http.NotFoundHandler()
			handler.ServeHTTP(w, r)
		}

	})
}

func main() {
	story := make(map[string]*Entry)

	jsonFile, err := os.Open("gopher.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	error := json.Unmarshal([]byte(byteValue), &story)

	if error != nil {
		fmt.Println(error)
	}
	htmlHander := HTMLHandler(story, "intro")
	log.Fatal(http.ListenAndServe(":8080", htmlHander))
}
