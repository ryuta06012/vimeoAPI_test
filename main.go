package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/ryuta06012/vimeoAPI_test/live"
	"github.com/ryuta06012/vimeoAPI_test/thumnail"
	"github.com/ryuta06012/vimeoAPI_test/videos"
)

func main() {
	addr := "localhost:4242"
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	http.HandleFunc("/upload_video", videos.UploadVideoHandler)
	http.HandleFunc("/get", thumnail.Get)
	http.HandleFunc("/post", thumnail.Post)
	http.HandleFunc("/live", live.CreateLivemovie)
	/* http.HandleFunc("/put", thumnail.Put)
	http.HandleFunc("/patch", thumnail.Patch) */

	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
