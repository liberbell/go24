package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func Home(w http.ResponseWriter, r *http.Request)  {
	data := 	
}