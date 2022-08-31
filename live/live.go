package live

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ryuta06012/vimeoAPI_test/request"
)

type Upload struct {
	Approach string `json:"approach,omitempty"`
}

type UploadRequest struct {
	Upload *Upload `json:"upload,omitempty"`
}

type CreateResponse struct {
	URI         string      `json:"uri"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	Link        string      `json:"link"`
}

type Status struct {
	Status string `json:"status"`
}

type PatchRequest struct {
	Live *Status `json:"live"`
}

type Live struct {
	ID                 int         `json:"id"`
	Link               string      `json:"link"`
	Key                string      `json:"key"`
	ActiveTime         interface{} `json:"active_time"`
	EndedTime          interface{} `json:"ended_time"`
	ArchivedTime       interface{} `json:"archived_time"`
	Status             string      `json:"status"`
	ScheduledStartTime interface{} `json:"scheduled_start_time"`
	SecondsRemaining   int         `json:"seconds_remaining"`
	SessionID          string      `json:"session_id"`
}

type PatchResponse struct {
	Live *Live `json:"live"`
}

func CreateLivemovie(w http.ResponseWriter, r *http.Request) {
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Content-Type":  "application/json",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	url := "https://api.vimeo.com/me/videos"
	var uploadRequest = &UploadRequest{
		Upload: &Upload{
			Approach: "live",
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
	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
	io.Copy(os.Stdout, res.Body)
	/* var live = &CreateResponse{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := json.Unmarshal(body, &live); err != nil {
		fmt.Println(err.Error())
	}
	patch(live) */
}

func patch(live *CreateResponse) {
	url := "https://api.vimeo.com" + live.URI
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Content-Type":  "application/json",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	livereq := &PatchRequest{
		Live: &Status{
			Status: "ready",
		},
	}
	req, err := request.NewRequest1("PATCH", url, header, livereq)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	io.Copy(os.Stdout, res.Body)
	/* var liveres = &PatchResponse{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := json.Unmarshal(body, &liveres); err != nil {
		fmt.Println(err.Error())
	} */
}

func Get(live *CreateResponse) {
	url := "https://api.vimeo.com" + live.URI
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("PREMIUM_ACCESS_TOKEN"),
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	req, err := request.NewRequest1("PATCH", url, header, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	io.Copy(os.Stdout, res.Body)
}
