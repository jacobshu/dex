package pokedex

type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Areas    []Area `json:"results"`
}
