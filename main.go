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
	client.Videos.UploadVideo("/Users/handaryuuta/Downloads/test1.mp4")
	/* //fmt.Printf("v: %v\n", v)
	_, thum, _ := client.Videos.GetThumnail(v.URI)
	fmt.Printf("thum.URI: %v\n", thum.URI)
	fmt.Printf("thum.Active: %v\n", thum.Active)
	fmt.Printf("thum.DefaultPicture: %v\n", thum.DefaultPicture)
	println("--------\n")
	client.Videos.UploadThumnail(v.URI, "/Users/handaryuuta/Downloads/test.png")
	client.Videos.UploadThumnail(v.URI, "/Users/handaryuuta/Downloads/pattern2_g.png")
	_, thum, _ = client.Videos.GetThumnail(v.URI)
	fmt.Printf("thum.URI: %v\n", thum.URI)
	fmt.Printf("thum.Active: %v\n", thum.Active) */
	//fmt.Printf("thum.DefaultPicture: %v\n", thum.DefaultPicture)
	//fmt.Printf("v: %v\n", v)
}
