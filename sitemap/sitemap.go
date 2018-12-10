package main

import (
	"encoding/xml"
	"flag"
	"os"
	"strings"

	"github.com/user/go_learning/link"
)

type urlbase struct {
	XMLName xml.Name `xml:"url"`
	URL     string   `xml:"loc"`
}
type urlset struct {
	XMLName xml.Name `xml:"urlset"`
	URLS    []urlbase
}

var webpagemap = make(map[string]bool)
var basesite string
var depth int

func main() {
	flag.StringVar(&basesite, "url", "text", "the url to search")
	flag.IntVar(&depth, "depth", 2, "how deep to go")
	flag.Parse()

	Parser(basesite, 0)
	f, err := os.Create("out.xml")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	var url urlset
	f.WriteString(xml.Header)
	for key := range webpagemap {
		testentry := urlbase{URL: key}
		url.URLS = append(url.URLS, testentry)
	}
	out, _ := xml.MarshalIndent(url, "", "	")
	f.WriteString(string(out))
}

// Parser grabs the url and determines whether we need to grab the links for it
func Parser(site string, level int) {
	if _, ok := webpagemap[site]; ok {
	} else {
		if level < depth {
			webpagemap[site] = true
			//if url is not already in our webpagemap. add it and follow down the rabbit hole
			CheckWebPageForLinks(site, level+1)
		}
	}
}

// CheckWebPageForLinks grabs any relevant links
func CheckWebPageForLinks(url string, level int) {
	webpage := link.ParseForALinks(url)
	if webpage != nil {
		for key := range webpage {
			if len(key) > 1 && key[0] == '/' && !strings.Contains(key, "www") && !strings.Contains(key, "com") { //we can assume this to be a local link
				//follow the link and add its values to our webpagemap
				Parser(basesite+key, level)
			} else if strings.Contains(key, basesite) {
				Parser(key, level)
			}
		}
	}
}
