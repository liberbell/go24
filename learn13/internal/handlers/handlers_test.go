package handlers

import "net/http"

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name   string
	url    string
	method string
	params []postData
	expectedStatusCode int
}{
	{"name", "/", "GET", []postData, http.StatusOK},
}

func TestHandlers(t *testing.T)  {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
}