package s3storage

import (
	"bytes"
	"context"
	"net/http"

	"multifinancetest/helpers/constants/rpcstd"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (storage *awsS3Implement) UploadFile(ctx context.Context, filename string, file []byte) (err error) {
	_, err = storage.client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(storage.AWS_BUCKET_NAME),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(file),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, "Failed to upload file to S3")
	}
	return nil
}
