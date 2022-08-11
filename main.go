package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"strconv"
	"encoding/json"

	"golang.org/x/oauth2"

	"github.com/joho/godotenv"
	"github.com/silentsokolov/go-vimeo/vimeo"
	tus "github.com/eventials/go-tus"
)

type Path struct {
	Filepath string `json:"filepath"`
}

type VideoID struct {
	ID string `json:"id"`
}

type Thumanil struct {
	Filepath string `json:"filepath"`
	ID string `json:"id"`
}

type Uploader struct{}

func Initialization() *vimeo.Client {
	//uploaderをクライアント構成に設定する
	config := vimeo.Config{
		Uploader: &Uploader{},
	}

	//トークンを使用して認証
	token := os.Getenv("ACCESS_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	//fmt.Println("ts = ", ts)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	// configを使用してvimeoのクライアントを作成
	client := vimeo.NewClient(tc, &config)

	return client
}

func RequestRead(w http.ResponseWriter, r *http.Request, method string, s interface{}) {
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    	return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
}
/*
	* videoのIDからvimeoにアップロードしたビデオを取得する
*/
func (v VideoID)GetVideo(c *vimeo.Client, vid int) (*vimeo.Video, error) {
	video, _, err := c.Videos.Get(vid)
	return video, err
}

func GetThumnailHandler(w http.ResponseWriter, r *http.Request) {
	var videoid  VideoID

	RequestRead(w, r, "GET", &videoid)
	client := Initialization()
	vid, _ := strconv.Atoi(videoid.ID)
	fmt.Println(vid)
	video, err := videoid.GetVideo(client, vid)
	if err != nil {
		log.Printf("videoid.GetVideo(): %v", err)
	}
	picture, _, err := client.Videos.GetPictures(video.GetID(), video.Pictures.GetID())
	if err != nil {
		log.Printf("Videos.GetPictures(): %v", err)
		return
	}
	fmt.Println("URI = ",picture.URI, "\n", "picture.Link = ", picture.Link)
}

func UploadThumnailHandler(w http.ResponseWriter, r *http.Request) {
	var thumnail Thumanil

	RequestRead(w, r, "GET", &thumnail)
	client := Initialization()
	vid, _ := strconv.Atoi(thumnail.ID)
	input := &vimeo.PicturesRequest{
		Active: true,
	}
	f, err := os.Open(thumnail.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	picture, _, err := client.Videos.CreatePictures(vid, input)
	if err != nil {
		log.Printf("Videos.CreatePictures(): %v", err)
		return
	}
	picture, _, err = client.Videos.UploadPicture(vid, input, f)
	if err != nil {
		log.Printf("Videos.UploadPicture(): %v", err)
		return
	}
	fmt.Println(picture)
}

func (u Uploader) UploadFromFile(c *vimeo.Client, uploadURL string, f *os.File) error {
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
// /videos/738517569
func UploadVideo(w http.ResponseWriter, filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	client := Initialization()
	//videoをアップロード
	video, _, _ := client.Users.UploadVideo("", f)
	
	fmt.Println("video.UR = ", video.URI, "\n", "video.Name = ", video.Name, "\n", "video.Link = ", video.Link, "\n", "video.Files = ", video.Files)
}

func UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	var path Path

	RequestRead(w, r, "POST", &path)
	if len(path.Filepath) < 1 {
		log.Printf("specify URI")
	}
	UploadVideo(w, path.Filepath)
}

func main() {
	addr := "localhost:4242"
	err := godotenv.Load(".env")
  	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	http.HandleFunc("/upload_video", UploadVideoHandler)
	http.HandleFunc("/get_thumnail", GetThumnailHandler)
	http.HandleFunc("/upload_thumnail", UploadThumnailHandler)
	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}