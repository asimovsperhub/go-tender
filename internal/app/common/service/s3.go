package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/internal/app/common/s3"
)

var S3 s3.S3Service

func InitializeS3(ctx g.Ctx) {
	S3 = s3.S3ToService(ctx)
}
