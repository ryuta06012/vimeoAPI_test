package vimeo

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	defaultBaseURL   = "https://api.vimeo.com"
	mediaTypeVersion = "application/vnd.vimeo.*+json;version=3.4"
)

type Service struct {
	client *Client
}

type Client struct {
	client *http.Client
	Token  string

	BaseURL *url.URL

	Videos *VideosService
}

type VimeoResponse struct {
	*http.Response
	// Pagination
	Page       int
	TotalPages int
	Total      int
	NextPage   string
	PrevPage   string
	FirstPage  string
	LastPage   string
}

func NewClient(httpClient *http.Client, accessTokn string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL, Token: accessTokn}
	c.Videos = &VideosService{client: c}
	return c
}

func (c *Client) Client() *http.Client {
	return c.client
}

func (c *Client) NewHttpRequest(method, url string, header map[string]string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	return req, nil
}

func (c *Client) NewFileRequest(method, url string, header map[string]string, body *os.File) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*VimeoResponse, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	response := newResponse(resp)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}
	return response, err
}

func newResponse(r *http.Response) *VimeoResponse {
	response := &VimeoResponse{Response: r}
	return response
}
