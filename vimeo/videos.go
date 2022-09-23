package vimeo

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	tus "github.com/eventials/go-tus"
)

type VideoUpload struct {
	Approach string `json:"approach,omitempty"`
	Size     int64  `json:"size,omitempty"`
}

type VideoUploadRequest struct {
	Upload *VideoUpload `json:"upload,omitempty"`
}

type VideoUploadResponse struct {
	URI  string
	HTML string
}

type VideosService Service

func uploadFromFile(uploadURL string, f *os.File) error {
	tusClient, err := tus.NewClient(uploadURL, nil)
	if err != nil {
		return err
	}

	upload, err := tus.NewUploadFromFile(f)
	if err != nil {
		return err
	}

	uploader := tus.NewUploader(tusClient, uploadURL, upload, 0)

	return uploader.Upload()
}

func (s *VideosService) GetVideo(vuri string) (*VimeoResponse, *UploadVideoResponse, error) {
	url := s.client.BaseURL.String() + vuri
	header := map[string]string{
		"Authorization": "bearer " + s.client.Token,
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	req, err := s.client.NewHttpRequest("GET", url, header, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	var videoRespponse = &UploadVideoResponse{}
	res, err := s.client.Do(req, videoRespponse)
	if err != nil {
		fmt.Println(err.Error())
	}
	return res, videoRespponse, err
}

func (s *VideosService) UploadVideo(videoUrl string) (*VideoUploadResponse, *VimeoResponse, error) {
	url := "https://api.vimeo.com/me/videos"
	header := map[string]string{
		"Authorization": "bearer " + s.client.Token,
		"Content-Type":  "application/json",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	file, _ := os.Open(videoUrl)
	stat, _ := file.Stat()
	if stat.IsDir() {
		return nil, nil, errors.New("the video file can't be a directory")
	}
	uploadRequest := &VideoUploadRequest{
		Upload: &VideoUpload{
			Approach: "tus",
			Size:     stat.Size(),
		},
	}
	req, err := s.client.NewHttpRequest("POST", url, header, uploadRequest)
	if err != nil {
		return nil, nil, err
	}
	var v = &UploadVideoResponse{}
	response, err := s.client.Do(req, v)
	if err != nil {
		return nil, response, nil
	}
	err = uploadFromFile(v.Upload.UploadLink, file)
	if err != nil {
		return nil, nil, err
	}
	vresp := NewVideoUploadReponse(v)
	return vresp, response, err
}

func (s *VideosService) VideoPatch(v *UploadVideoResponse) {
	header := map[string]string{
		"Authorization": "bearer " + s.client.Token,
		"Tus-Resumable": "1.0.0",
		"Upload-Offset": "0",
		"Content-Type":  "application/offset+octet-stream",
	}
	req, err := s.client.NewHttpRequest("PATCH", v.Upload.UploadLink, header, nil)
	_, err = s.client.Do(req, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	s.VideoHead(v)
}

func (s *VideosService) VideoHead(v *UploadVideoResponse) {
	header := map[string]string{
		"Tus-Resumable": "1.0.0",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	req, err := s.client.NewHttpRequest("HEAD", v.Upload.UploadLink, header, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
}

func NewVideoUploadReponse(v *UploadVideoResponse) *VideoUploadResponse {
	var vresp = &VideoUploadResponse{}

	vresp.URI = v.URI
	vresp.HTML = v.Embed.HTML
	return vresp
}
