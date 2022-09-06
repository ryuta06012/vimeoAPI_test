package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/ryuta06012/vimeoAPI_test/vimeo"
)

// "/Users/handaryuuta/Downloads/test.png"
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	client := vimeo.NewClient(nil, "264de4b79b29af50f49579bb8d12283c")
	v, _, _ := client.Videos.UploadVideo("/Users/handaryuuta/Downloads/test1.mp4")
	fmt.Printf("v: %v\n", v)
	client.Videos.UploadThumnail(v.URI, "/Users/handaryuuta/Downloads/test.png")
	client.Videos.UploadThumnail(v.URI, "/Users/handaryuuta/Downloads/pattern2_g.png")
	fmt.Printf("v: %v\n", v)
}
