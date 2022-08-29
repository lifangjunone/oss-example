package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"oss-example/conf"
)

var (
	client   *oss.Client
	err      error
	filePath string
)

func init() {
	client, err = oss.New(conf.OssEndPoint, conf.OssAccessKeyId, conf.OssAccessSecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func validate() error {
	if conf.OssAccessSecret == "" || conf.OssAccessKeyId == "" || conf.OssAccessSecret == "" {
		return fmt.Errorf("oss connection params is empty")
	}
	if filePath == "" {
		return fmt.Errorf("file is not exist")
	}
	return nil
}

func loadParams() {
	flag.StringVar(&filePath, "f", "files/下载1.jpeg", "Please entry file path")
	flag.Parse()
}

func upload(filePath string) error {
	bucket, err := client.Bucket(conf.OssBucket)
	if err != nil {
		return err
	}
	return bucket.PutObjectFromFile(filePath, filePath)
}

func main() {
	loadParams()
	if err := validate(); err != nil {
		fmt.Printf("params is missing")
		os.Exit(1)
	}
	if err := upload(filePath); err != nil {
		fmt.Printf("upload file is error : %s", err)
		os.Exit(1)
	}
}
