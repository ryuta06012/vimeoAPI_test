package auth

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ryuta06012/vimeoAPI_test/views"
	"golang.org/x/oauth2"
)

func Tinitialization1() string {
	url := "https://api.vimeo.com/oauth/authorize/client"

	client := &http.Client{}

	// ヘッダー作成
	header := http.Header{}
	token := "basic " + b64.StdEncoding.EncodeToString([]byte(os.Getenv("CLIENTID")+":"+os.Getenv("CLIENTSECRET")))
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", token)
	header.Add("Accept", "application/vnd.vimeo.*+json;version=3.4")

	// ボディ作成
	reqBody := views.RequestToken{
		Grante_type: "client_credentials",
		Scope:       "private edit upload video_files public",
	}
	jsonValue, _ := json.Marshal(reqBody)
	//reqbody := []byte(`{"grant_type": "client_credentials", "scope": "private edit upload video_files public"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer((jsonValue)))
	req.Header = header
	res, _ := client.Do(req)
	var response views.Response
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	fmt.Println("Token = ", response.AccessToken, "response.Scope = ", response.Scope)
	io.Copy(os.Stdout, res.Body)
	return response.AccessToken
}

func Tinitialization() *http.Client {
	token := os.Getenv("PREMIUM_ACCESS_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return tc
}
