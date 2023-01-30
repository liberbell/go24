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

	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
		} else {

		}
	}
}