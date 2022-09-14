package vimeo

import (
	"errors"
	"fmt"
	"os"
)

type PathRequest struct {
	Active bool `json:"active"`
}

/*
Get the thumbnail set for the video
*/
func (s *VideosService) GetThumnail(vuri string) (*Response, *ThumnailReponse, error) {
	_, video, err := s.GetVideo(vuri)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	fmt.Printf("video.Pictures.URI: %v\n", video.Pictures.URI)
	res, thumnail, err := getThumnail(s.client, video.Pictures.URI)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	return res, thumnail, err
}

func (s *VideosService) UploadThumnail(vuri string, filpath string) (*Response, *ThumnailReponse, error) {
	url := s.client.BaseURL.String() + vuri + "/pictures"
	_, thumanilreponse, err := createThumnailResource(s.client, url)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}

	_, err = uploadThumnai(s.client, thumanilreponse.Link, filpath)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}

	response, thumnailres, err := editThumnai(s.client, thumanilreponse.URI)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	return response, thumnailres, nil
}

func createThumnailResource(c *Client, url string) (*Response, *ThumnailReponse, error) {
	header := map[string]string{
		"Authorization": "bearer " + c.Token,
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	body := []byte("{\"time\": 100000}")
	req, err := c.NewRequest1("POST", url, header, body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var thumanilresponse = &ThumnailReponse{}
	res, err := c.Do(req, thumanilresponse)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	return res, thumanilresponse, nil
}

func uploadThumnai(c *Client, url string, filpath string) (*Response, error) {
	header := map[string]string{
		"Authorization": "bearer " + c.Token,
		"Content-Type":  "image/png",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	file, err := os.Open(filpath)
	req, err := c.NewRequest2("PUT", url, header, file)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	response, err := c.Do(req, nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return response, nil
}

func editThumnai(c *Client, uri string) (*Response, *ThumnailReponse, error) {
	url := c.BaseURL.String() + uri
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
		return nil, nil, errors.New(err.Error())
	}
	var thumanilresponse = &ThumnailReponse{}
	reponse, err := c.Do(req, thumanilresponse)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	return reponse, thumanilresponse, nil
}

func getThumnail(c *Client, uri string) (*Response, *ThumnailReponse, error) {
	url := c.BaseURL.String() + uri
	header := map[string]string{
		"Authorization": "bearer " + c.Token,
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	req, err := c.NewRequest1("GET", url, header, nil)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	var thumnailresponse = &ThumnailReponse{}
	reponse, err := c.Do(req, thumnailresponse)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	return reponse, thumnailresponse, nil
}
