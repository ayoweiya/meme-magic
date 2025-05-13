package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

const (
	MinioEndpoint  = "localhost:9000"
	MinioAccessKey = "admin"
	MinioSecretKey = "password123"
	MinioBucket    = "meme-magic"
)

func ConnectMinIO() (*minio.Client, error) {
	minioClient, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
		Secure: false, // 如果是 HTTPS 則改成 true
	})

	if err != nil {
		log.Println("❌ 無法連接 MinIO:", err)
		return nil, err
	}

	// 確保 bucket 存在
	exists, err := minioClient.BucketExists(context.Background(), MinioBucket) // 建立一個沒有超時、取消的 context
	if err != nil {
		return nil, err
	}
	if !exists {
		err = minioClient.MakeBucket(context.Background(), MinioBucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("❌ 建立 MinIO 儲存桶失敗: %v", err)
		}
		log.Println("✅ 已建立 MinIO 儲存桶:", MinioBucket)
	}

	return minioClient, nil
}

// 上傳圖片到 MinIO
func UploadToMinIO(fileName string, fileData []byte) (string, error) {
	client, err := ConnectMinIO()
	if err != nil {
		return "", err
	}

	// 上傳圖片
	_, err = client.PutObject(context.Background(), MinioBucket, fileName,
		bytes.NewReader(fileData), int64(len(fileData)), minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		return "", fmt.Errorf("❌ 上傳 MinIO 失敗: %v", err)
	}

	// 回傳可存取的 URL
	imageURL := fmt.Sprintf("http://%s/%s/%s", MinioEndpoint, MinioBucket, fileName)
	return imageURL, nil
}
