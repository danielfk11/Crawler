package main

import (
	"danielfk11/website"
	"danielfk11/crawler"
	"flag"
	"fmt"
)

var(
 link string
 action string
)

func init()  {
	flag.StringVar(&link, "url", "https://aprendagolang.com.br", "url para inicar busca")
	flag.StringVar(&action, "action", "website", "qual servico iniciar")
}

func main() {
	flag.Parse()

	switch *&action{
	case "website":
		website.Run()

	case "webcrawler":
	
	done := make(chan bool)
	wb := crawler.New()
	go wb.VisitLink(link)

	for log := range wb.Log() {
		fmt.Println(log)
	}

	<-done

	default:
		fmt.Printf("action '%s' nao foi reconhecida", action)
	}
}
