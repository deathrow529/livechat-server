package awswrapper

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewAWSSession : Creates new AWS session
func NewAWSSession(region string) *session.Session {
	config := &aws.Config{
		Region: aws.String(region),
	}
	return session.Must(session.NewSession(config))
}

// NewDynamoDBService : Creates new Dynamodb instance
func NewDynamoDBService(sess *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(sess)
}

// UpdateItem : Update Dynamodb Item
func UpdateItem(svc *dynamodb.DynamoDB, tableName string,
	key map[string]*dynamodb.AttributeValue,
	updateMap map[string]*dynamodb.AttributeValueUpdate) {
	input := &dynamodb.UpdateItemInput{
		TableName:        aws.String(tableName),
		Key:              key,
		AttributeUpdates: updateMap,
	}
	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
