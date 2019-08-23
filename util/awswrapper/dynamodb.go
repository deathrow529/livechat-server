package awswrapper

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamodbUpdateItem : Update Dynamodb Item
func DynamodbUpdateItem(svc *dynamodb.DynamoDB, tableName string,
	key map[string]*dynamodb.AttributeValue,
	updateMap map[string]*dynamodb.AttributeValueUpdate) {
	fmt.Println(tableName)
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
