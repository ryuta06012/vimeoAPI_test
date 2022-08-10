package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"encoding/json"

	"golang.org/x/oauth2"

	"github.com/joho/godotenv"
	"github.com/silentsokolov/go-vimeo/vimeo"
	tus "github.com/eventials/go-tus"
)

type Path struct {
	Filepath string `json:"filepath"`
}

type Uploader struct{}

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
func UploadVideoHandler(w http.ResponseWriter, filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	//uploaderをクライアント構成に設定する
	config := vimeo.Config{
		Uploader: &Uploader{},
	}

	//トークンを使用して認証
	token := os.Getenv(".env")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	//fmt.Println("ts = ", ts)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	// configを使用してvimeoのクライアントを作成
	client := vimeo.NewClient(tc, &config)

	//videoをアップロード
	video, _, _ := client.Users.UploadVideo("", f)
	
	fmt.Println(video.URI)
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	var path Path

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    	return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&path); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	if len(path.Filepath) < 1 {
		log.Printf("specify URI")
	}
	UploadVideoHandler(w, path.Filepath)
}

func main() {
	addr := "localhost:4242"
	err := godotenv.Load(".env")
  	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	http.HandleFunc("/test", TestFunc)
	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}