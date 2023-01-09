package main

import (
	"fmt"
	"io/ioutil"

	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	len, err := file.WriteString(" Creating first s3 bucket")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
}
func ReadFile() {
	fmt.Printf("\n\nReading a file in Go lang\n")
	fileName := "test.txt"
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
}
func UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	return err
}
func main() {
	CreateFile()
	ReadFile()
	sess, err1 := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},
	})
	if err1 != nil {
		fmt.Printf("Failed to initialize new session: %v", err1)
		return
	}
	bucketName := "golangshashi"
	uploader := s3manager.NewUploader(sess)
	filename := "test.txt"
	err := UploadFile(uploader, "test.txt", bucketName, filename)
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}
	fmt.Println("Successfully uploaded file!")
}
