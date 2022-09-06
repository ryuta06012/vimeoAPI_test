package vimeo

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type PathRequest struct {
	Active bool `json:"active"`
}

func (s *VideosService) ThumnailGet(vuri string) {
	url := s.client.BaseURL.String() + vuri + "?fields=metadata.connections.pictures.uri"
	header := map[string]string{
		"Authorization": "bearer " + s.client.Token,
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	req, err := s.client.NewRequest1("GET", url, header, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	io.Copy(os.Stdout, res.Body)
}

func (s *VideosService) UploadThumnail(vuri string, purl string) error {
	url := s.client.BaseURL.String() + vuri + "/pictures"
	header := map[string]string{
		"Authorization": "bearer " + s.client.Token,
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	body := []byte("{\"time\": 100000}")
	req, err := s.client.NewRequest1("POST", url, header, body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var response = &PostResponse{}
	_, err = s.client.Do(req, response)
	if err != nil {
		fmt.Println(err.Error())
	}
	uploadThumnai(s.client, response.Link, purl, response.URI)
	return nil
}

func uploadThumnai(c *Client, url string, purl string, picture_uri string) {
	header := map[string]string{
		"Authorization": "bearer " + c.Token,
		"Content-Type":  "image/png",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	file, err := os.Open(purl)
	req, err := c.NewRequest2("PUT", url, header, file)
	if err != nil {
		fmt.Println("NewRequest1 = ", err.Error())
	}
	_, err = c.Do(req, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	editThumnai(c, picture_uri)
}

func editThumnai(c *Client, uri string) {
	url := "https://api.vimeo.com" + uri
	header := map[string]string{
		"Authorization": "bearer " + c.Token,
		"Content-Type":  "application/json",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	body := &PathRequest{
		Active: true,
	}
	req, err := c.NewRequest1("PATCH", url, header, body)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = c.Do(req, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
