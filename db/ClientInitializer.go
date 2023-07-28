package db

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DynamoDbClient struct {
	Client    *dynamodb.Client
	TableName string
}
