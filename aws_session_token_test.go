package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"testing"
)

func TestGetAWSSessionToken(t *testing.T) {
	conf := &aws.Config{
		Region:      aws.String(""),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	}

	sess, err := session.NewSession(conf)
	if err != nil {
		t.Fatal(err)
	}
	svc := sts.New(sess)
	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(3600),
		SerialNumber:    aws.String(""),
		TokenCode:       aws.String(""),
	}

	result, err := svc.GetSessionToken(input)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
