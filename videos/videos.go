package videos

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	tus "github.com/eventials/go-tus"
	"github.com/ryuta06012/vimeoAPI_test/request"
	"github.com/ryuta06012/vimeoAPI_test/views"
)

type Upload struct {
	Approach string `json:"approach,omitempty"`
	Size     int64  `json:"size,omitempty"`
}

type UploadRequest struct {
	Upload *Upload `json:"upload,omitempty"`
}

func UploadFromFile(uploadURL string, f *os.File) error {
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

func UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.vimeo.com/me/videos" //POST
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Content-Type":  "application/json",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	file, _ := os.Open("/Users/handaryuuta/Downloads/test_movie.mp4")
	stat, _ := file.Stat()
	fmt.Printf("stat.Size(): %v\n", stat.Size())
	uploadRequest := &UploadRequest{
		Upload: &Upload{
			Approach: "tus",
			Size:     stat.Size(),
		},
	}
	req, err := request.NewRequest1("POST", url, header, uploadRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	//io.Copy(os.Stdout, res.Body)
	defer res.Body.Close()

	var v = &views.UploadVideoResponse{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := json.Unmarshal(body, &v); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("v.Upload.UploadLink: %v\n", v.Upload.UploadLink)
	fmt.Printf("v.Link: %v\n", v.Link)
	fmt.Printf("v.Approach: %v\n", v.Upload.Approach)
	fmt.Printf("v.Pictures.URI: %v\n", v.Pictures.URI)
	fmt.Printf("v.URI: %v\n", v.URI)
	err = UploadFromFile(v.Upload.UploadLink, file)
	if err != nil {
		fmt.Println(err.Error())
	}
	//Patch(v)
}

func Patch(v *views.UploadVideoResponse) {
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Tus-Resumable": "1.0.0",
		"Upload-Offset": "0",
		"Content-Type":  "application/offset+octet-stream",
	}
	file, _ := os.Open("/Users/handaryuuta/Downloads/test_movie.mp4")
	req, err := request.NewRequest1("PATCH", v.Upload.UploadLink, header, file)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)
	//io.Copy(os.Stdout, resp.Body)
	fmt.Println("resp: = ", resp.Header["Upload-Offset"])
	fmt.Println("resp: = ", resp.Header["Upload-Length"])
	//io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
	Head(v)
}

func Head(v *views.UploadVideoResponse) {
	header := map[string]string{
		"Tus-Resumable": "1.0.0",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	//file, _ := os.Open("/Users/handaryuuta/Downloads/test_movie.mp4")
	req, err := request.NewRequest1("HEAD", v.Upload.UploadLink, header, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("resp: = ", resp.Header.Get("Upload-Offset"))
	fmt.Println("resp: = ", resp.Header["Upload-Length"])
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)
	//io.Copy(os.Stdout, resp.Body)
	fmt.Println("resp: = ", resp.Header["Upload-Offset"])
	fmt.Println("resp: = ", resp.Header["Upload-Length"])
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}
