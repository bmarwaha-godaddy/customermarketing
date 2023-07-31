package db

import (
	"CustomerMarketingPlatform/initializer"
	"CustomerMarketingPlatform/model"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

func StoreData(basics initializer.DynamoDbClient, channel model.Channel) error {

	item, err := attributevalue.MarshalMap(channel)

	if err != nil {
		log.Fatalf("Error parsing customer channel %v", err)
	}

	_, err = basics.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: item,
	})

	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}

func ReadDataById(basics initializer.DynamoDbClient, identifier string, loggedInCity string) (error, model.Channel) {
	var channel model.Channel
	id, err := attributevalue.Marshal(identifier)
	loggedCityMarshal, err := attributevalue.Marshal(loggedInCity)

	response, err := basics.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(basics.TableName),
		Key:       map[string]types.AttributeValue{"identifier": id, "loggedInFrom": loggedCityMarshal},
	})

	if err != nil {
		log.Printf("Couldn't get item from table. Here's why: %v\n", err)
		return err, channel
	}

	err = attributevalue.UnmarshalMap(response.Item, &channel)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
	}

	return err, channel
}
func DeleteChannelById(basics initializer.DynamoDbClient, identifier string, loggedInCity string) error {
	id, err := attributevalue.Marshal(identifier)
	loggedCityMarshal, err := attributevalue.Marshal(loggedInCity)

	response, err := basics.Client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		Key:       map[string]types.AttributeValue{"identifier": id, "loggedInFrom": loggedCityMarshal},
		TableName: aws.String(basics.TableName),
	})

	log.Printf("DeleteResponse  is %v", response)

	if err != nil {
		log.Printf("Couldn't get item from table. Here's why: %v\n", err)
		return err
	}

	return err
}
