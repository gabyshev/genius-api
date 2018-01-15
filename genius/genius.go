package genius

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	//baseURL is the endpoint for all API methods
	baseURL string = "https://api.genius.com"
)

//Client is a client for Genius API
type Client struct {
	AccessToken string
	client      *http.Client
}

//NewClient creates Client to work with Genius API
//You can pass http.Client or it will use http.DefaultClient by default
//
//It requires a token for accessing Genius API
func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{AccessToken: token, client: httpClient}
	return c
}

//doRequest makes a request and puts authorization token in headers
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

//GetAccount returns current user account data
func (c *Client) GetAccount() (*User, error) {
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

	if response.Meta.Status != 200 {
		return nil, errors.New(response.Meta.Message)
	}

	return response.Response.User, nil
}

//GetArtist returns Artist object in response
//
//Uses "dom" as textFormat by default
func (c *Client) GetArtist(id int) (*Artist, error) {
	return c.GetArtistDom(id)
}

//GetArtistDom returns Artist object in response
//With "dom" as textFormat
func (c *Client) GetArtistDom(id int) (*Artist, error) {
	return c.getArtist(id, "dom")
}

//GetArtistPlain returns Artist object in response
//With "plain" as textFormat
func (c *Client) GetArtistPlain(id int) (*Artist, error) {
	return c.getArtist(id, "plain")
}

//GetArtistHTML returns Artist object in response
//With "html" as textFormat
func (c *Client) GetArtistHTML(id int) (*Artist, error) {
	return c.getArtist(id, "html")
}

//GetArtistSongs returns array of songs objects in response
func (c *Client) GetArtistSongs(id int, sort string, per_page int, page int) ([]*Song, error) {
	url := fmt.Sprintf(baseURL+"/artists/%d/songs", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("sort", sort)
	q.Add("per_page", string(per_page))
	q.Add("page", string(page))
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

	if response.Meta.Status != 200 {
		return nil, errors.New(response.Meta.Message)
	}

	return response.Response.Songs, nil
}

//GetSong returns Song object in response
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

//getArtist is a method taking id and textFormat as arguments to make request and return Artist object in response
func (c *Client) getArtist(id int, textFormat string) (*Artist, error) {
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

	if response.Meta.Status != 200 {
		return nil, errors.New(response.Meta.Message)
	}

	return response.Response.Artist, nil
}

//Search returns array of Hit objects in response
//
//Currently only songs are searchable by this handler
func (c *Client) Search(q string) ([]*Hit, error) {
	url := fmt.Sprintf(baseURL + "/search")
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	getParams := req.URL.Query()
	getParams.Add("q", q)
	req.URL.RawQuery = getParams.Encode()

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Meta.Status != 200 {
		return nil, errors.New(response.Meta.Message)
	}

	return response.Response.Hits, nil
}

//GetAnnotation gets annotation object in response
func (c *Client) GetAnnotation(id int, textFormat string) (*Annotation, error) {
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

	response.Response.Annotation.Process(textFormat)

	if response.Meta.Status != 200 {
		return nil, errors.New(response.Meta.Message)
	}

	return response.Response.Annotation, nil
}
