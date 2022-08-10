# vimeoAPI_test
## Running the sample
1. Run the server
~~~
go run server.go
~~~
2. Send a POST request
~~~
curl -H "content-type: application/json" -X POST -d'{"filepath":"mp4ファイルの絶対パスを指定"}' http://localhost:4242/test
~~~
