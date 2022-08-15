package minio

import (
	"context"
	"drive/config"
	"drive/tools"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"sort"
	"strings"
	"sync"
)

type Client struct {
	client *minio.Client
}

func NewMinioClient() *Client {

	minioClient, err := minio.New(config.Conf.MinioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.MinioConfig.AccessKey, config.Conf.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logrus.Errorf("fail creating minioClient: %v", err)
		return nil
	}

	return &Client{client: minioClient}
}

var defaultPutObjectOptions minio.PutObjectOptions

func (c *Client) UploadFiles(ctx context.Context, files []*multipart.FileHeader) ([]*UploadResult, error) {
	var res []*UploadResult

	wg := &sync.WaitGroup{}
	concurrentLimiter := make(chan bool, 10)
	defer close(concurrentLimiter)

	wgResponseChan := &sync.WaitGroup{}
	responseChan := make(chan *UploadResult, 5)
	go func() {
		wgResponseChan.Add(1)
		for response := range responseChan {
			res = append(res, response)
		}
		// 手动关闭 responseChan 才会运行到这
		wgResponseChan.Done()
	}()

	for index, file := range files {
		suffix := file.Filename[strings.LastIndex(file.Filename, "."):]
		objectName := fmt.Sprintf("%d%s", tools.GenSnowflakeID(), suffix)
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		wg.Add(1)
		concurrentLimiter <- true
		go c.PutOneFile(ctx, index, f, objectName, file.Size, responseChan, concurrentLimiter, wg)
	}

	wg.Wait()
	// wg.Wait 使 responseChan 确定不需要再接收了，才可以 close
	close(responseChan)
	wgResponseChan.Wait()

	// 子任务错误校验
	for _, singleRes := range res {
		if singleRes.Err != nil {
			return nil, singleRes.Err
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Index < res[j].Index
	})

	return res, nil
}

type UploadResult struct {
	Index int    `json:"index"`
	Url   string `json:"url"`
	Err   error  `json:"err,omitempty"`
}

func (c *Client) PutOneFile(ctx context.Context, index int, file io.Reader, objectName string, fileSize int64, responseChan chan *UploadResult, concurrentLimiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		_, err := c.client.PutObject(ctx, "drive", objectName, file, fileSize, defaultPutObjectOptions)
		if err != nil {
			logrus.Errorf("GoRoutine ERROR! minio putObject error: %v", err)
		}
	}()

	responseChan <- &UploadResult{
		Index: index,
		Url:   fmt.Sprintf("http://%s/drive/%s", config.Conf.MinioConfig.Endpoint, objectName),
		Err:   nil,
	}
	<-concurrentLimiter
}

func (c *Client) ListFiles(ctx context.Context) []string {
	objChan := c.client.ListObjects(ctx, "drive", minio.ListObjectsOptions{
		WithVersions: false,
		WithMetadata: false,
		Prefix:       "",
		Recursive:    false,
		MaxKeys:      0,
		StartAfter:   "",
		UseV1:        false,
	})

	res := make([]string, 0)
	for obj := range objChan {
		res = append(res, obj.Key)
	}
	return res
}