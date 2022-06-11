package website


import (
	"fmt"
	"danielfk11/database"
	"net/http"
	"html/template"

)


type DataLinks struct {
	Links []database.VisitedLink
}


func indexHandle() func(http.ResponseWriter, *http.Request) {
	tmpl, err := template.ParseFiles("website/templates/index.html", "website/templates/style.css")
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		links, err := database.FindAllLinks()
		if err != nil {

		}

		fmt.Println(links)

		data := DataLinks{Links: links}

		tmpl.Execute(w, data)
	}
}