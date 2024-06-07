package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"net/http"
)

type S3Servant struct {
	client *minio.Client
}

func (s *S3Servant) S3Upload(fd []byte, filetype string, filename string, filesize int64) string {
	// objectsize 传递'-1'将占用内存
	tmp := filetype + "/" + filename
	info, err := s.client.PutObject(context.Background(), "tender", tmp, bytes.NewReader(fd), filesize,
		minio.PutObjectOptions{ContentType: http.DetectContentType(fd)})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("INFO------------------------------>", info.Location)
	return info.Location
}

func (s *S3Servant) ExistBucket(bulkName string) bool {
	exist, err := s.client.BucketExists(context.Background(), bulkName)
	if err != nil {
		log.Fatalf("is exist bulk eror " + err.Error())
		return false
	}
	return exist
}

func (s *S3Servant) GetObject(bulkName string, ObjectName string) *minio.Object {
	object, err := s.client.GetObject(context.Background(), bulkName, ObjectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
	}
	//defer object.Close()
	return object
}
