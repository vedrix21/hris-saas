package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"context"
	"io"

	"github.com/google/uuid"
)

const maxFileSize = 2 << 20 // 2MB

func SaveUpload(file *multipart.FileHeader, folder string) (string, error) {
	// 🔒 Validasi size
	if file.Size > maxFileSize {
		return "", errors.New("file too large (max 2MB)")
	}

	// 🔒 Validasi extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".pdf" {
		return "", errors.New("invalid file type (only jpg, png, pdf)")
	}

	// 🔒 Validasi MIME type (anti file palsu)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buffer)

	if filetype != "image/jpeg" &&
		filetype != "image/png" &&
		filetype != "application/pdf" {
		return "", errors.New("invalid file content")
	}

	// 🔥 Generate nama unik
	filename := fmt.Sprintf("%d_%s%s",
		time.Now().Unix(),
		uuid.New().String(),
		ext,
	)

	// 📁 Path final
	dir := filepath.Join("uploads", folder)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(dir, filename)

	return fullPath, nil
}

func NewS3Client() *s3.Client {
	endpoint := os.Getenv("AWS_ENDPOINT")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
	)
	if err != nil {
		panic(err)
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
}

func UploadToS3(file io.Reader, filename string, contentType string) (string, error) {
	client := NewS3Client()
	bucket := os.Getenv("AWS_BUCKET")

	key := fmt.Sprintf("payments/%d_%s", time.Now().Unix(), filename)

	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", err
	}

	return key, nil
}

func GeneratePresignedURL(key string) (string, error) {
	client := NewS3Client()
	bucket := os.Getenv("AWS_BUCKET")

	presignClient := s3.NewPresignClient(client)

	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		return "", err
	}

	return req.URL, nil
}