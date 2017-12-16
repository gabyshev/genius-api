package genius

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.genius.com"

type Client struct {
	AccessToken string
	client      *http.Client
}

func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{AccessToken: token, client: httpClient}
	return c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func (c *Client) GetAccount() (*Response, error) {
	url := fmt.Sprintf(baseURL + "/account/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetArtist(id int) (*Response, error) {
	return c.getArtist(id, "dom")
}

func (c *Client) GetArtistDom(id int) (*Response, error) {
	return c.getArtist(id, "dom")
}

func (c *Client) GetArtistPlain(id int) (*Response, error) {
	return c.getArtist(id, "plain")
}

func (c *Client) GetArtistHTML(id int) (*Response, error) {
	return c.getArtist(id, "html")
}

func (c *Client) GetArtistSongs(id int, sort string, per_page int, page int) {
	url := fmt.Sprintf(baseURL+"/artists/%d/songs", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("sort", sort)
	q.Add("per_page", per_page)
	q.Add("page", page)
	req.URL.RawQuery = q.Encode()

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	response.Response.Artist.Process(textFormat)

	return &response, nil
}

func (c *Client) GetSong(id int, textFormat string) (*Response, error) {
	url := fmt.Sprintf(baseURL+"/songs/%d", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("text_format", textFormat)
	req.URL.RawQuery = q.Encode()

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	response.Response.Song.Process(textFormat)

	return &response, nil

}

func (c *Client) getArtist(id int, textFormat string) (*Response, error) {
	url := fmt.Sprintf(baseURL+"/artists/%d", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("text_format", textFormat)
	req.URL.RawQuery = q.Encode()

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	response.Response.Artist.Process(textFormat)

	return &response, nil

}

func (c *Client) Search(q string) (*Response, error) {
	url := fmt.Sprintf(baseURL + "/search")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", q)
	req.URL.RawQuery = q.Encode()

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	response.Response.Artist.Process(textFormat)

	return &response, nil

}

func (c *Client) GetAnnotation(id string, textFormat string) (*Response, error) {
	url := fmt.Sprintf(baseURL+"/annotations/%d", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("text_format", textFormat)
	req.URL.RawQuery = q.Encode()

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	response.Response.Artist.Process(textFormat)

	return &response, nil

}
