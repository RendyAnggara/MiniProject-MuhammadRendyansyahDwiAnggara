package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v3"
)

func SaveUploadedFile(c fiber.Ctx, file *multipart.FileHeader) (string, error) {
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	savePath := filepath.Join("./uploads", filename)

	if err := c.SaveFile(file, savePath); err != nil {
		return "", err
	}

	return filename, nil
}