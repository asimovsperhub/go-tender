package s3

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"sync"
)

var onceS3 sync.Once
var S3 S3Service

func S3ToService(ctx g.Ctx) S3Service {
	onceS3.Do(func() {
		S3 = NewS3Service(ctx)
	})
	return S3
}

type S3Service interface {
	S3Upload(fd []byte, filetype string, filename string, filesize int64) string
	ExistBucket(bulkName string) bool
	GetObject(bulkName string, ObjectName string) *minio.Object
}

func NewS3Service(ctx g.Ctx) S3Service {
	// s := g.Cfg()
	//42.193.247.183
	endpoint := "127.0.0.1:9000"
	accessKeyID := "JvTpb5Gnt4qCL69v"
	secretAccessKey := "AnyVSSQK2aTWM2QmE3HjZQB1040gYB5L"
	useSSL := false
	// Initialize minio client object.
	minioClient, _ := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	s3 := &S3Servant{
		client: minioClient,
	}
	return s3
}
