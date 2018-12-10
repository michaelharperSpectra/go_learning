package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/user/go_learning/link"
)

var webpagemap = make(map[string]struct{})
var basesite string
var exists = struct{}{}

func main() {
	flag.StringVar(&basesite, "url", "text", "the url to search")
	flag.Parse()

	Parser(basesite)
	fmt.Println(webpagemap)
}

// Parser grabs the url and determines whether we need to grab the links for it
func Parser(site string) {
	fmt.Println(site)
	if _, ok := webpagemap[site]; ok {
		fmt.Println(1)
	} else {
		fmt.Println(2)
		//if url is not already in our webpagemap. add it and follow down the rabbit hole
		CheckWebPageForLinks(site)
		webpagemap[site] = exists

	}

}

// CheckWebPageForLinks grabs any relevant links
func CheckWebPageForLinks(url string) {
	webpage := link.ParseForALinks(url)

	for key := range webpage {
		fmt.Println(key)
		if key[0] == '/' && !strings.Contains(key, "www") && !strings.Contains(key, "com") { //we can assume this to be a local link
			//follow the link and add its values to our webpagemap
			Parser(basesite + key)
		} else if strings.Contains(key, basesite) {
			Parser(key)
		}
	}
}
