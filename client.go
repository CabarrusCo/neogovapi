package neogovapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	username   string
	password   string
	httpClient *http.Client
}

func NewClient(username string, password string) *Client {
	return &Client{
		username:   username,
		password:   password,
		httpClient: &http.Client{},
	}
}

func (c *Client) Send(r *http.Request, v interface{}) error {
	r.Header.Set("Accept", "application/json")

	if len(r.Header.Get("Content-Type")) == 0 {
		r.Header.Set("Content-Type", "application/json")
	}

	r.SetBasicAuth(c.username, c.password)

	res, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("Error encountered on %v. Got status code %v", r.URL, res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(v)

}
