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

var testCollection = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "last_name", value: "john@smith.com"},
		{key: "phone", value: "000 0000 0000"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {

	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, test := range testCollection {
		if test.method == "GET" {
			resp, err := testServer.Client().Get(testServer.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range test.params {
				values.Add(x.key, x.value)
			}
			resp, err := testServer.Client().PostForm(testServer.URL+test.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
