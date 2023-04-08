package upsdk

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	Accounts     *AccountsService
	Categories   *CategoriesService
	Utils        *UtilsService
	Tags         *TagsService
	Transactions *TransactionsService
}

type transport struct {
	underlyingTransport http.RoundTripper
	token               string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return t.underlyingTransport.RoundTrip(req)
}

func NewClient(httpClient *http.Client, token string) *Client {
	baseURL, _ := url.Parse("https://api.up.com.au/api/v1")
	httpClient.Transport = &transport{
		underlyingTransport: http.DefaultTransport,
		token:               token,
	}

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}

	c.Accounts = &AccountsService{c}
	c.Categories = &CategoriesService{c}
	c.Utils = &UtilsService{c}
	c.Tags = &TagsService{c}
	c.Transactions = &TransactionsService{c}

	return c
}

func (c *Client) NewRequest(method, url string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.BaseURL, url), nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

type PaginationLinks struct {
	Previous string `json:"prev"`
	Next     string `json:"next"`
}

type ResponseMessage struct {
	Body       interface{}     `json:"data"`
	Pagination PaginationLinks `json:"links"`
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	defer io.Copy(ioutil.Discard, res.Body)

	response := &Response{Response: res}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(v)
		}
	}

	return response, err
}

type Response struct {
	*http.Response

	NextPageUrl string
	PrevPageUrl string
}
