package conf

import (
	"os"
)

var (
	OssAccessKeyId  string = os.Getenv("OssAccessKeyId")
	OssAccessSecret string = os.Getenv("OssAccessSecret")
	OssEndPoint     string = os.Getenv("OssEndPoint")
	OssBucket       string = os.Getenv("OssBucket")
)

func init() {
	OssAccessKeyId = os.Getenv("OssAccessKeyId")
	OssAccessSecret = os.Getenv("OssAccessSecret")
	OssEndPoint = os.Getenv("OssEndPoint")
	OssBucket = os.Getenv("OssBucket")
}
