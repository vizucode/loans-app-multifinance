package s3storage

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/vizucode/gokit/utils/env"
)

type awsS3Implement struct {
	client          *s3.S3
	AWS_REGION      string
	AWS_SECRET_KEY  string
	AWS_ACCESS_KEY  string
	AWS_BUCKET_NAME string
}

func NewAwsS3Implement(
	AWS_REGION string,
	AWS_SECRET_KEY string,
	AWS_ACCESS_KEY string,
	AWS_BUCKET_NAME string,
	AWS_ENDPOINT string,
) *awsS3Implement {

	creds := credentials.NewStaticCredentials(
		AWS_ACCESS_KEY,
		AWS_SECRET_KEY,
		"",
	)

	awsConfig := &aws.Config{
		Region:      aws.String(AWS_REGION), // Sesuaikan region langsung di sini
		Credentials: creds,
	}

	if env.GetBool("AWS_S3_FORCE_PATH_STYLE") {
		awsConfig.S3ForcePathStyle = aws.Bool(true)
	}

	if !strings.EqualFold(AWS_ENDPOINT, "") {
		awsConfig.Endpoint = aws.String(AWS_ENDPOINT)
	}

	// start session s3
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &awsS3Implement{
		client:          s3.New(sess),
		AWS_REGION:      AWS_REGION,
		AWS_SECRET_KEY:  AWS_SECRET_KEY,
		AWS_ACCESS_KEY:  AWS_ACCESS_KEY,
		AWS_BUCKET_NAME: AWS_BUCKET_NAME,
	}
}
