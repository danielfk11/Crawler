package crawler

import (
	"danielfk11/database"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
)

type webCrawler struct {
	log chan string
}

func (wb *webCrawler)VisitLink(link string)  {
	
  	wb.log <- fmt.Sprintf("visitando: %s", link)
	fmt.Println(link)

	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	wb.log <- fmt.Sprintf("[ERRO DETECTADO] status diferente de 200: %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	wb.extractLinks(doc)
	
}

func (wb *webCrawler)extractLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href"{
				continue
			}
			
			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" || link.Scheme == "mailto" || link.Scheme == "javascript" || link.Scheme == "tel" || link.Scheme == "itmss"{
				continue
			}

			if database.CheckVisitedLink(link.String()) {
				wb.log <- fmt.Sprintf("link ja visitado: %s", link)
				continue
			}

			visitedLink := database.VisitedLink{
				Website: link.Host,
				Link: link.String(),
				VisitedDate: time.Now(),
			}

			database.Insert("links", visitedLink)

		//	links = append(links, link.String()) 

			go wb.VisitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		wb.extractLinks(c)
	}
}

func (wb *webCrawler) Log() (chan string) {
	return wb.log
}

func New() *webCrawler  {
	wb := &webCrawler{
		log: make(chan string, 10),
	}
	return wb
}