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
	help     bool = false
)

// init
func init() {
	client, err = oss.New(conf.OssEndPoint, conf.OssAccessKeyId, conf.OssAccessSecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

// validate
func validate() error {
	if conf.OssAccessSecret == "" || conf.OssAccessKeyId == "" || conf.OssAccessSecret == "" {
		return fmt.Errorf("oss connection params is empty")
	}
	if filePath == "" {
		return fmt.Errorf("file is not exist")
	}
	return nil
}

// loadParams
func loadParams() {
	flag.StringVar(&filePath, "f", "files/下载1.jpeg", "Please entry file path")
	flag.BoolVar(&help, "h", false, "Print help info")
	flag.Parse()
	if help {
		usage()
	}
}

// print usage info
func usage() {
	// print project description info
	fmt.Fprintf(os.Stderr, `
		name: oss-example
		version: v1.0.0
		Usage: oss-example [-h] -f <upload file path>
		Options: 
	`)
	// print could usage params
	flag.PrintDefaults()
}

// upload file
func upload(filePath string) error {
	bucket, err := client.Bucket(conf.OssBucket)
	if err != nil {
		return err
	}
	return bucket.PutObjectFromFile(filePath, filePath)
}

// download file
func download(filePath string) error {
	bucket, err := client.Bucket(conf.OssBucket)
	downloadURL, err := bucket.SignURL(filePath, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("file download url: %s", downloadURL)
	return nil
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
