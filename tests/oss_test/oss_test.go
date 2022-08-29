package oss_test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"oss-example/conf"
	"path"
	"testing"
)

var (
	client *oss.Client
	err    error
)

func init() {
	client, err = oss.New(conf.OssEndPoint, conf.OssAccessKeyId, conf.OssAccessSecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func TestOssPackage(t *testing.T) {
	t.Logf("Oss version: %s", oss.Version)
}

// TestUploadFile
func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(conf.OssBucket)
	if err != nil {
		t.Logf("error: %s", err)
	}
	basePath, err := os.Getwd()
	if err != nil {
		t.Logf("get current file path is error: %s", err)
	}
	filePath := path.Join(basePath, "../../files/下载4.jpeg")
	err = bucket.PutObjectFromFile("下载4.jpeg", filePath)
	if err != nil {
		t.Logf("upload file is error: %s", err)
	}
}

// TestBucketList
func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Logf("list bucket is error: %s", err)
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("Bucket:", bucket.Name)
	}
}
