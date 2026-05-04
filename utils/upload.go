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