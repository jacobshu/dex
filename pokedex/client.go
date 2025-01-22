package pokedex

import (
	"fmt"
	"net/http"
)

func makeReq(endpoint string, params map[string]string) (*http.Request, error) {
	baseurl := "https://pokeapi.co/api/v2"
	fullurl := baseurl + endpoint

	req, err := http.NewRequest(http.MethodGet, fullurl, nil)
	if err != nil {
		return &http.Request{}, fmt.Errorf("CommandMap: error creating request: %w\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}
