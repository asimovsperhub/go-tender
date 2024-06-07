package qbucket

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"io/ioutil"
	"log"
)

type CosServant struct {
	client *cos.Client
}

func (s *CosServant) Upload(fd io.Reader, filetype string, filename string, filesize int64) error {
	tmp := filetype + "/" + filename
	_, err := s.client.Object.Put(context.Background(), tmp, fd, nil)
	log.Println(err)
	if err != nil {
		log.Println("Cos Upload err------------->", err)
	}
	return nil
}

func (s *CosServant) GetObject(ObjectName string) []byte {
	// name := "ygp/1.png"
	resp, err := s.client.Object.Get(context.Background(), ObjectName, nil)
	if err != nil {
		log.Println("Cos GetObject err------------->", err)
		// panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return bs
}
