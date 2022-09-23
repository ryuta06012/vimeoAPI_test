package movie

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/ryuta06012/vimeoAPI_test/vimeo"
)

func GetMovie() error {
	sess := createSession()

	bucketName := "dev-musemo-movie-bucket"
	objectKey := "movie/te.mp4"
	/* var name = flag.String("name", "John", "User name.")
	flag.Parse() */
	filename := "sample.mp4"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	defer f.Close()
	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	uploadVideo(filename)
	println("success!\n")
	return nil
}

func uploadVideo(url string) error {
	client := vimeo.NewClient(nil, os.Getenv("PREMIUM_ACCESS_TOKEN"))
	uploadRes, _, err := client.Videos.UploadVideo(url)
	if err != nil {
		return err
	}
	fmt.Printf("uploadRes.HTML: %v\n", uploadRes.HTML)
	return nil
}

func createSession() *session.Session {
	// 特に設定しなくても環境変数にセットしたクレデンシャル情報を利用して接続してくれる
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
		Region:      aws.String("ap-northeast-1"),
		//s3://dev-musemo-movie-bucket/movie/
		//Endpoint:         aws.String("http://minio:9000"), // コンテナ内からアクセスする場合はホストをサービス名で指定
		//S3ForcePathStyle: aws.Bool(true), // ローカルで動かす場合は必須
	}
	return session.Must(session.NewSession(&cfg))
}

func existsBucket(svc *s3.S3, bucket string) error {
	_, err := svc.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	return nil
}
