package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"

	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint        string
	accessKeyID     string
	accessKeySecret string
	osspath         string
	bucketName      string
)

func init() {
	// 配置
	endpoint = os.Getenv("endpoint")
	if endpoint == "" {
		log.Fatal("Repo: Settings->Secrets need ENDPOINT")
	}
	accessKeyID = os.Getenv("access_key_id")
	if accessKeyID == "" {
		log.Fatal("Repo: Settings->Secrets need ACCESS_KEY_ID")
	}
	accessKeySecret = os.Getenv("access_key_secret")
	if accessKeySecret == "" {
		log.Fatal("Repo: Settings->Secrets need ACCESS_KEY_SECRET")
	}
	osspath = os.Getenv("osspath")
	if osspath == "" {
		log.Fatal("Repo: Settings->Secrets need OSSPATH")
	}
	bucketName = os.Getenv("bucket_name")
	if bucketName == "" {
		log.Fatal("Repo: Settings->Secrets need BUCKET_NAME")
	}
}

func main() {
	f, err := os.Open("url")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		url := string(a)
		if url != "" {
			downloadToAlioss(url)
		}
	}
}

func downloadToAlioss(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Get url success:", url)

	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get oss client success.")

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get oss bucket success.")

	filePath := path.Join(osspath, path.Base(url))
	err = bucket.PutObject(filePath, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Put file to bucket success. File path:", filePath)
}
