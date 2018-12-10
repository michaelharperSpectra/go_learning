package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/user/go_learning/link"
)

var webpagemap = make(map[string]bool)
var basesite string

func main() {
	flag.StringVar(&basesite, "url", "text", "the url to search")
	flag.Parse()

	Parser(basesite, 0)
	fmt.Println(webpagemap)
}

// Parser grabs the url and determines whether we need to grab the links for it
func Parser(site string, level int) {
	if _, ok := webpagemap[site]; ok {
	} else {
		if level < 2 {
			fmt.Println(site)
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
			if len(key) > 0 && key[0] == '/' && !strings.Contains(key, "www") && !strings.Contains(key, "com") { //we can assume this to be a local link
				//follow the link and add its values to our webpagemap
				Parser(basesite+key, level)
			} else if strings.Contains(key, basesite) {
				Parser(key, level)
			}
		}
	}
}
