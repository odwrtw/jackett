package jackett

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Client represents a jackett client
type Client struct {
	URL    string
	APIKey string
}

// New returns a new jackett client
func New(url, apiKey string) *Client {
	return &Client{
		URL:    url,
		APIKey: apiKey,
	}
}

// Indexer represents an indexer
type Indexer struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Status  int    `json:"Status"`
	Results int    `json:"Results"`
	Error   string `json:"Error"`
}

// Result represents the jackett results
type Result struct {
	Tracker      string `json:"Tracker"`
	TrackerID    string `json:"TrackerId"`
	CategoryDesc string `json:"CategoryDesc"`
	Title        string `json:"Title"`
	GUID         string `json:"Guid"`
	Link         string `json:"Link"`
	Size         int    `json:"Size"`
	Seeders      int    `json:"Seeders"`
	Peers        int    `json:"Peers"`
}

// Response represents a response
type Response struct {
	Results  []*Result  `json:"Results"`
	Indexers []*Indexer `json:"Indexers"`
}

// Search searches for stuff
func (c *Client) Search(query string, trackers []Tracker, categories []Category) (*Response, error) {
	v := url.Values{}
	v.Add("apikey", c.APIKey)
	v.Add("Query", query)
	for _, c := range categories {
		v.Add("Category", strconv.Itoa(int(c)))
	}
	for _, t := range trackers {
		v.Add("Tracker", string(t))
	}

	url := c.URL + "/api/v2.0/indexers/all/results?" + v.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &Response{}
	return r, json.NewDecoder(resp.Body).Decode(r)
}
