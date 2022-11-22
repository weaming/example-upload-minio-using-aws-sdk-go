package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("example.jpg")
	if err != nil {
		panic(err)
	}
	cfg := &S3Config{
		AccessKeyID:     "ZSCR0B4UJ2ZRQWAWUX0Z",
		AccessKeySecret: "I9PDoAzKuTAQsHCpRvsC4jEId8jwsnmvBfIWVMVq",
		Endpoint:        "http://localhost:9000",
		Region:          "ap-east-1",
		BucketName:      "test",
	}
	url, err := UploadObject("/img/example.jpg", dat, cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}
