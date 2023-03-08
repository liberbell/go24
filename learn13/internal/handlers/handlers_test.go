package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	// {"name", "/", "GET", []postData{}, http.StatusOK},
	// {"about", "/about", "GET", []postData{}, http.StatusOK},
	// {"quarters", "/general-quarters", "GET", []postData{}, http.StatusOK},
	// {"major", "/majors-suite", "GET", []postData{}, http.StatusOK},
	// {"search", "/search-availability", "GET", []postData{}, http.StatusOK},
	// {"contact", "/contact", "GET", []postData{}, http.StatusOK},
	// {"makereservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	// {"postavailability", "/search-availability", "POST", []postData{
	// 	{key: "start", value: "2023-01-31"},
	// 	{key: "end", value: "2023-01-31"},
	// }, http.StatusOK},
	// {"postavailabilityjson", "/search-availability-json", "POST", []postData{
	// 	{key: "start", value: "2023-01-31"},
	// 	{key: "end", value: "2023-01-31"},
	// }, http.StatusOK},
	// {"makereservation", "/make-reservation", "POST", []postData{
	// 	{key: "first_name", value: "eric"},
	// 	{key: "last_name", value: "clapton"},
	// 	{key: "email", value: "a@b.com"},
	// 	{key: "phone", value: "123-456-7890"},
	// }, http.StatusOK},
}

func TestHandlers(t *testing.T) {
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
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General Quarters",
		},
	}

}
