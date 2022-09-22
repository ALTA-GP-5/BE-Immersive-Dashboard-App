package helpers

import (
	"errors"
	"immersive/config"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToS3(directory string, fileName string, contentType string, fileData multipart.File) (string, error) {

	// The session the S3 Uploader will use
	sess := config.GetSession()

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String("/" + directory + "/" + fileName),
		Body:        fileData,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", errors.New(err.Error())
	}
	return result.Location, nil
}

func CheckFileExtension(filename string) (string, error) {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])

	if extension != "jpg" && extension != "jpeg" && extension != "png" && extension != "pdf" {
		return "", errors.New("forbidden file type")
	}

	return extension, nil
}

func CheckFileSize(size int64, contentType string) error {
	if size == 0 {
		return errors.New("illegal file size")
	}

	if contentType != "pdf" {
		if size > 10097152 {
			return errors.New("file size too big")
		}
	}

	if size > 1097152 {
		return errors.New("file size too big")
	}

	return nil
}

func UploadPDFToS3(directory string, fileName string, contentType string, data io.Reader) (string, error) {

	// The session the S3 Uploader will use
	sess := config.GetSession()

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String("/" + directory + "/" + fileName),
		Body:        data,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", errors.New(err.Error())
	}
	return result.Location, nil
}
