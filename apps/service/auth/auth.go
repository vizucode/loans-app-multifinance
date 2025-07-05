package auth

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"multifinancetest/apps/repositories"
	"multifinancetest/helpers/constants/rpcstd"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nfnt/resize"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

type auth struct {
	db        repositories.IDatabase
	fs        repositories.IStorage
	validator *validator.Validate
}

func NewAuth(
	db repositories.IDatabase,
	fs repositories.IStorage,
	validator *validator.Validate,
) *auth {
	return &auth{
		db:        db,
		fs:        fs,
		validator: validator,
	}
}

func (uc *auth) processImage(ctx context.Context, base64Str, dimensions string) ([]byte, error) {
	const maxFileSize = 5 * 1024 * 1024 // 5MB
	base64Len := len(base64Str)
	fileSize := (base64Len * 3 / 4) - strings.Count(base64Str, "=") // Hitung ukuran file asli

	if fileSize > maxFileSize {
		err := errorkit.NewErrorStd(http.StatusRequestEntityTooLarge, rpcstd.INVALID_ARGUMENT, "file size exceeds the 5MB limit")
		logger.Log.Error(ctx, err)
		return nil, err
	}
	// Decode base64 string menjadi byte array
	imageData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		logger.Log.Error(ctx, err)
		err := errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.INVALID_ARGUMENT, "invalid base64 string")
		return nil, err
	}

	// Deteksi tipe file berdasarkan header
	fileType := http.DetectContentType(imageData[:512])
	if fileType != "image/jpeg" && fileType != "image/png" {
		err := errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.INVALID_ARGUMENT, "file type must be jpg or png")
		logger.Log.Error(ctx, err)
		return nil, err
	}

	// Decode byte array menjadi image.Image
	var (
		img    image.Image
		width  uint
		height uint
	)

	switch fileType {
	case "image/jpeg":
		img, err = jpeg.Decode(bytes.NewReader(imageData))
	case "image/png":
		img, err = png.Decode(bytes.NewReader(imageData))
	default:
		return nil, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.INVALID_ARGUMENT, "file type must be jpg or png")
	}
	if err != nil {
		return nil, errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, "internal server error")
	}

	// Parsing dimension input
	if dimensions != "" {
		width, height = uint(0), uint(0)
		parts := strings.Split(dimensions, "x")
		if len(parts) == 2 {
			w, _ := strconv.Atoi(parts[0])
			h, _ := strconv.Atoi(parts[1])
			width, height = uint(w), uint(h)
		}
		if width == 0 || height == 0 {
			return nil, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.INVALID_ARGUMENT, "invalid dimensions. Format must be WIDTHxHEIGHT")
		}
	}
	// Resize gambar
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Encode ulang gambar setelah resize
	var buf bytes.Buffer
	switch fileType {
	case "image/jpeg":
		err = jpeg.Encode(&buf, resizedImg, &jpeg.Options{Quality: 70})
	case "image/png":
		err = png.Encode(&buf, resizedImg)
	}
	if err != nil {
		logger.Log.Error(ctx, err)
		return nil, errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, "internal server error")
	}

	return buf.Bytes(), nil
}
