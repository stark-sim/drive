package pkg

import (
	"context"
	"drive/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"io"
)

func NewMinioClient() *minio.Client {

	minioClient, err := minio.New(config.Conf.MinioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.MinioConfig.AccessKey, config.Conf.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logrus.Errorf("fail creating minioClient: %v", err)
		return nil
	}

	return minioClient
}

func UploadFiles(ctx context.Context, minioClient *minio.Client, file io.Reader, fileSize int64) {
	_, err := minioClient.PutObject(ctx, "drive", "fileName", file, fileSize, minio.PutObjectOptions{})
	if err != nil {
		return
	}
}
