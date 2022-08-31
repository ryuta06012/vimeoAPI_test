package thumnail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ryuta06012/vimeoAPI_test/request"
	"github.com/ryuta06012/vimeoAPI_test/views"
)

func Get(w http.ResponseWriter, r *http.Request) {
	//token := auth.Tinitialization1()
	url := "https://api.vimeo.com/videos/743575151?fields=metadata.connections.pictures.uri"
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	req, err := request.NewRequest1("GET", url, header, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	io.Copy(os.Stdout, res.Body)
}

func Post(w http.ResponseWriter, r *http.Request) {
	//token := auth.Tinitialization1()
	url := "https://api.vimeo.com/videos/744418128/pictures" //POST
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	body := []byte("{\"time\": 100000}")
	req, err := request.NewRequest1("POST", url, header, body)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	var response views.PostResponse
	body, _ = io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	fmt.Println("URI = ", response.URI, "\n", "Link = ", response.Link)
	//io.Copy(os.Stdout, res.Body)
	Put(w, r, response.Link, response.URI)
}

func Put(w http.ResponseWriter, r *http.Request, url string, picture_uri string) {
	//client := auth.Tinitialization()
	//url := "https://kaiju.cloud.vimeo.com/video/1495831421?expires=1661548095&sig=7570b8285c322b95a4f4addde68668a564b674a0"

	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Content-Type":  "image/png",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	file, err := os.Open("/Users/handaryuuta/Downloads/test.png")
	/* wf := multipart.NewWriter(&buf)
	f, err := os.Open("/Users/handaryuuta/Downloads/test.png")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	fw, err := wf.CreateFormFile("file", "/Users/handaryuuta/Downloads/test.png")
	if err != nil {
		fmt.Println("CreateFormFile = ", err.Error())
	}
	if _, err = io.Copy(fw, f); err != nil {
		fmt.Println("Copy = ", err.Error())
	}
	wf.Close() */
	req, err := request.NewRequest2("PUT", url, header, file)
	if err != nil {
		fmt.Println("NewRequest1 = ", err.Error())
	}
	//req.Header.Set("Content-Type", wf.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	io.Copy(os.Stdout, res.Body)
	fmt.Printf("picture_uri: %v\n", picture_uri)
	Patch(w, r, picture_uri)
}

func Patch(w http.ResponseWriter, r *http.Request, uri string) {
	//client := auth.Tinitialization()
	url := "https://api.vimeo.com" + uri //PATCH
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Content-Type":  "application/json",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	body := &views.PathRequest{
		Active: true,
	}
	req, err := request.NewRequest1("PATCH", url, header, body)
	fmt.Println("-------------\n")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("-------------\n")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("res: %v\n", res.Body)
	io.Copy(os.Stdout, res.Body)
}
