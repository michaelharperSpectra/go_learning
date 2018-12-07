package link

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// ParseForALinks parse an HTML file and extract all of the <a> links
// 	filename: string parameter for the file to be parsed
//  returns map of the href value to all text stored inside <a></a> tags
//  returns nil if invalid html
func ParseForALinks(filename string) map[string][]string {
	var r io.Reader
	r, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	z := html.NewTokenizer(r)

	var isInsideA bool
	var currentlink string
	var allText []string
	var returnval = make(map[string][]string)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return returnval
			}
			return nil
		case html.TextToken:
			if isInsideA {
				s := string(z.Text())
				allText = append(allText, s)
			}
		case html.StartTagToken, html.EndTagToken:

			token := z.Token()
			if token.Data == "a" {
				if tt == html.StartTagToken {
					isInsideA = true
					for _, attr := range token.Attr {
						if attr.Key == "href" {
							currentlink = attr.Val
						}
					}
				} else {
					returnval[currentlink] = allText
					fmt.Println(currentlink)
					fmt.Println(returnval[currentlink])
					isInsideA = false
					currentlink = ""
					allText = nil
				}
			}
		}

	}
}
