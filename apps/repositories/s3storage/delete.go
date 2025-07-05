package s3storage

import (
	"context"
	"net/http"
	"strings"

	"multifinancetest/helpers/constants/rpcstd"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (storage *awsS3Implement) DeleteFile(ctx context.Context, key string) error {
	if strings.TrimSpace(key) == "" {
		return nil
	}

	// Kirim permintaan ke S3 untuk menghapus file
	_, err := storage.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(storage.AWS_BUCKET_NAME),
		Key:    aws.String(key),
	})
	if err != nil {
		logger.Log.Error(ctx, "Error deleting file from S3: ", err)
		return errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, "Failed to delete file from S3")
	}

	return nil
}
