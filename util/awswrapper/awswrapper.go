package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewAWSSession : Creates new AWS session
func NewAWSSession(region string) *session.Session {
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	return session
}

// NewDynamoDBService : Creates new Dynamodb instance
func NewDynamoDBService(sess *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(sess)
}
