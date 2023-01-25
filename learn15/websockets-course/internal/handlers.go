package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(jet.VarMap)
	data.Set("user_id", 1)
	data.Set("first", "Eric")
	dow := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	data.Set("dow", dow)
	err := renderPage(w, "home.jet", data)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error executing template: ", err)
	}
}

func SendData(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "send.jet", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error executing template: ", err)
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println("Unexpected template err: ", err)
	}
}
