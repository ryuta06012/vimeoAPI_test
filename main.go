package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/oauth2"

	b64 "encoding/base64"

	tus "github.com/eventials/go-tus"

	"github.com/joho/godotenv"
	"github.com/silentsokolov/go-vimeo/vimeo"
)

type Path struct {
	Filepath string `json:"filepath"`
}

type VideoID struct {
	ID string `json:"id"`
}

type Thumanil struct {
	Filepath string `json:"filepath"`
	ID       string `json:"id"`
}

type TpictureSize struct {
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	Link               string `json:"link,omitempty"`
	LinkWithPlayButton string `json:"link_with_play_button,omitempty"`
}

type Tpictures struct {
	URI            string          `json:"uri,omitempty"`
	Active         bool            `json:"active"`
	BaseLink       string          `json:"base_link"`
	DefaultPicture bool            `json:"default_picture"`
	Type           string          `json:"type,omitempty"`
	Sizes          []*TpictureSize `json:"sizes,omitempty"`
	Link           string          `json:"link,omitempty"`
	ResourceKey    string          `json:"resource_key,omitempty"`
}

type Request struct {
	Name string `json:"name"`
}

type Uploader struct{}

func Initialization() *vimeo.Client {
	//uploaderをクライアント構成に設定する
	config := vimeo.Config{
		Uploader: &Uploader{},
	}

	//トークンを使用して認証
	token := os.Getenv("PREMIUM_ACCESS_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
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
func (v VideoID) GetVideo(c *vimeo.Client, vid int) (*vimeo.Video, error) {
	video, _, err := c.Videos.Get(vid)
	return video, err
}

func GetThumnailHandler(w http.ResponseWriter, r *http.Request) {
	var videoid VideoID

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
	fmt.Println("URI = ", picture.URI, "\n", "picture.Link = ", picture.Link)
}

func NewRequest1(method, url string, header map[string]string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	// req, err := http.NewRequest(method, url, bytes.NewReader(body))
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	return req, nil
}
func NewRequest2(method, url string, header map[string]string, body *os.File) (*http.Request, error) {
	//var buf io.ReadWriter
	/* if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	} */
	// req, err := http.NewRequest(method, url, bytes.NewReader(body))
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	return req, nil
}

func Tinitialization() *http.Client {
	token := os.Getenv("PREMIUM_ACCESS_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return tc
}

type Response struct {
	BaseLink    string `json:"base_link"`
	Link        string `json:"link"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type RequestToken struct {
	Grante_type string `json:"grant_type"`
	Scope       string `json:"scope"`
}

func Tinitialization1() string {
	url := "https://api.vimeo.com/oauth/authorize/client"

	client := &http.Client{}

	// ヘッダー作成
	header := http.Header{}
	token := "Basic " + b64.StdEncoding.EncodeToString([]byte(os.Getenv("CLIENTID")+":"+os.Getenv("CLIENTSECRET")))
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", token)
	header.Add("Accept", "application/vnd.vimeo.*+json;version=3.4")

	// ボディ作成
	reqBody := RequestToken{
		Grante_type: "client_credentials",
		Scope:       "private, edit, upload, video_files, public",
	}
	jsonValue, _ := json.Marshal(reqBody)
	json.Unmarshal(jsonValue, &reqBody)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header = header
	res, _ := client.Do(req)
	var response Response
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	fmt.Println("Token = ", response.AccessToken, "response.Scope = ", response.Scope)
	io.Copy(os.Stdout, res.Body)
	return response.AccessToken
}

func Get(w http.ResponseWriter, r *http.Request) {
	token := Tinitialization1()
	url := "https://api.vimeo.com/videos/740217970?fields=metadata.connections.pictures.uri"
	header := map[string]string{
		"Authorization": "bearer " + token,
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	req, err := NewRequest1("GET", url, header, nil)
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
	token := Tinitialization1()
	url := "https://api.vimeo.com/videos/740217970/pictures" //POST
	header := map[string]string{
		"Authorization": "bearer " + token,
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	body := []byte("{\"time\": 100000}")
	req, err := NewRequest1("POST", url, header, body)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	/* var response Response
	body, _ = io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	fmt.Println("BaseLink = ", response.BaseLink, "\n", "Link = ", response.Link) */
	io.Copy(os.Stdout, res.Body)
	//Put(w, r, response.BaseLink)
}

func Put(w http.ResponseWriter, r *http.Request) {
	client := Tinitialization()
	//url := u
	url := "https://i.vimeocdn.com/video/1489716249-891aaf766698d4e072020e06deb479289701ebbf64e19f2d982f380d8d022205-d"
	header := map[string]string{
		"Accept":       "application/vnd.vimeo.*+json;version=3.4",
		"Content-Type": "image/png",
	}

	file, err := os.Open("/Users/handaryuuta/Downloads/test.png")
	/* img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Println("unable to encode image.")
	}
	imageBytes := buffer.Bytes() */
	req, err := NewRequest2("PUT", url, header, file)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	io.Copy(os.Stdout, res.Body)
}

func Patch(w http.ResponseWriter, r *http.Request) {
	client := Tinitialization()
	url := "https://api.vimeo.com/videos/739601966/pictures/1488476988" //PATCH
	header := map[string]string{
		"Authorization": "bearer " + os.Getenv("ACCESS_TOKEN"),
		"Content-Type":  "application/json",
		"Accept":        "application/vnd.vimeo.*+json;version=3.4",
	}
	body := []byte(`{"active": true}`)
	req, err := NewRequest1("PATCH", url, header, body)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	io.Copy(os.Stdout, res.Body)
}

func UploadThumnailHandler(w http.ResponseWriter, r *http.Request) {
	var thumnail Thumanil
	Initialization()
	RequestRead(w, r, "POST", &thumnail)
	client := Initialization()
	vid, _ := strconv.Atoi(thumnail.ID)
	input := &vimeo.PicturesRequest{
		Time:   float32(10000),
		Active: true,
	}
	f, err := os.Open(thumnail.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	/* picture, _, err := client.Videos.CreatePictures(vid, input)
	if err != nil {
		log.Printf("Videos.CreatePictures(): %v", err)
		return
	} */
	fmt.Println("id =", thumnail.ID, " filepath =", thumnail.Filepath)
	picture, _, err := client.Videos.UploadPicture(vid, input, f)
	if err != nil {
		log.Printf("Videos.UploadPicture(): %v", err)
		return
	}
	fmt.Println(picture, "\n")
	/* picture, _, err := client.Videos.EditPictures(vid, 1486761507, input)
	if err != nil {
		log.Printf("Videos.EditPictures(): %v", err)
		return
	}
	fmt.Println(picture, "\n") */
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

// /videos/738843229
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
	http.HandleFunc("/get", Get)
	http.HandleFunc("/post", Post)
	http.HandleFunc("/put", Put)
	http.HandleFunc("/patch", Patch)

	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
