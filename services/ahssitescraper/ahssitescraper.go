package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

const url string = "https://atlantahumane.org"

func main() {

	r, err := http.Get(url)
	if err != nil {
		println("Cannot get page")
	}

	doc, err := html.Parse(r.Body)
	if err != nil {
		println("Cannot parse page")
	}

	runs := 0
	var f func(*html.Node)
	f = func(n *html.Node) {
		if runs < 100 {

			if n.Type == html.ElementNode {
				// Do something with n...
				runs++
				fmt.Println(n.Type)
				fmt.Println(n.Data)
				fmt.Println(n.Attr)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(doc)

}
