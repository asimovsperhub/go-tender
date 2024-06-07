package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func Test_bucket(t *testing.T) {
	endpoint := "42.193.247.183:9000"
	accessKeyID := "JvTpb5Gnt4qCL69v"
	secretAccessKey := "AnyVSSQK2aTWM2QmE3HjZQB1040gYB5L"
	useSSL := false
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}
func Test_Check(t *testing.T) {
	endpoint := "42.193.247.183:9000"
	accessKeyID := "asimov"
	secretAccessKey := "asimov@123"
	useSSL := true
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup
	found, err := minioClient.BucketExists(context.Background(), "tender")
	if err != nil {
		fmt.Println(err)
		return
	}
	if found {
		fmt.Println("Bucket found")
	}
}

func existBulck(minioClient *minio.Client, bulkName string) bool {
	bool, err := minioClient.BucketExists(context.Background(), bulkName)
	if err != nil {
		log.Fatalf("is exist bulk eror " + err.Error())
		return false
	}
	return bool
}
func Test_s3(t *testing.T) {
	endpoint := "42.193.247.183:9000"
	accessKeyID := "asimov"
	secretAccessKey := "asimov@123"
	useSSL := false
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//exist := existBulck(minioClient, "tender")
	//fmt.Println("EXIST-------->", exist)
	//log.Printf("%#v\n", minioClient) // minioClient is now setup
	object, err := minioClient.GetObject(context.Background(), "tender", "/default/1.png", minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer object.Close()
	localFile, err := os.Create("local-file.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer localFile.Close()

	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
}

func Test_PutFile(t *testing.T) {
	f, err := os.Open("/Users/apple/Desktop/1.png")
	if err != nil {
		log.Println("read file fail", err)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	log.Println("fd--------------------->", bytes.NewReader(fd))
	if err != nil {
		log.Println("read to fd fail", err)
	}

	endpoint := "42.193.247.183:9000"
	accessKeyID := "JvTpb5Gnt4qCL69v"
	secretAccessKey := "AnyVSSQK2aTWM2QmE3HjZQB1040gYB5L"
	useSSL := false
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	// onjectsize 传递'-1'将占用内存
	tmp := "icon/3.png"
	info, err := minioClient.PutObject(context.Background(), "tender", tmp, bytes.NewReader(fd), -1,
		minio.PutObjectOptions{ContentType: http.DetectContentType(fd)})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("-------------------->", int64(len(fd)))
	log.Println("------------------->", info)
	// 图片链接
	log.Println("------------------->", info.Location)
}

func Test_Getbject(t *testing.T) {
	object := S3.GetObject("tender", "/default/1.png")
	fmt.Println(object)
}
