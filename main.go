package main

import (
	"danielfk11/crawler"
	"flag"
	"fmt"
)

var(
 link string
 action string
)

func init()  {
	flag.StringVar(&link, "url", "https://github.com/", "url para inicar busca")
}

func main() {
	flag.Parse()

	
	done := make(chan bool)
	wb := crawler.New()
	go wb.VisitLink(link)

	for log := range wb.Log() {
		fmt.Println(log)
	}

	<-done

}
