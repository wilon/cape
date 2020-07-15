package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cheggaaa/pb/v3"
)

var (
	osspath string
	client  *oss.Client
	bucket  *oss.Bucket
)

func init() {
	// 配置
	endpoint := os.Getenv("endpoint")
	if endpoint == "" {
		log.Fatal("Repo: Settings->Secrets need ENDPOINT")
	}
	accessKeyID := os.Getenv("access_key_id")
	if accessKeyID == "" {
		log.Fatal("Repo: Settings->Secrets need ACCESS_KEY_ID")
	}
	accessKeySecret := os.Getenv("access_key_secret")
	if accessKeySecret == "" {
		log.Fatal("Repo: Settings->Secrets need ACCESS_KEY_SECRET")
	}
	osspath = os.Getenv("osspath")
	if osspath == "" {
		log.Fatal("Repo: Settings->Secrets need OSSPATH")
	}
	bucketName := os.Getenv("bucket_name")
	if bucketName == "" {
		log.Fatal("Repo: Settings->Secrets need BUCKET_NAME")
	}
	// oss
	var err error
	client, err = oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get oss client success.")

	bucket, err = client.Bucket(bucketName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get oss bucket success.")
}

func main() {
	// 读取url文件
	f, err := os.Open("url")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	// 下载
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
	// get
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Get url success:", url, resp.ContentLength)

	// bar
	bar := pb.New64(resp.ContentLength)
	bar.Set(pb.Bytes, true)
	bar.SetRefreshRate(time.Second * time.Duration(resp.ContentLength/1024/1024/10))
	bar.SetTemplateString(`Downloding... {{counters . }} {{bar . }} {{percent . }} {{etime . "%s"}} {{speed . }} ` + "\n")
	bar.Start()
	defer bar.Finish()
	reader := bar.NewProxyReader(resp.Body)

	// put
	filePath := path.Join(osspath, path.Base(url))
	err = bucket.PutObject(filePath, reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Put file to bucket success. File path:", filePath)
}
