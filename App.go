package main

import (
	"CustomerMarketingPlatform/db"
	"CustomerMarketingPlatform/initializer"
	"CustomerMarketingPlatform/model"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"net/http"

	"context"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/channel", addChannelAccess)
	r.GET("/channel/", getChannelById)
	r.DELETE("/channel/", deleteChannelById)

	r.Run()

}
func getChannelById(c *gin.Context) {

	cfg, error := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))

	if error != nil {
		log.Fatalf("error acessing client %v", error)
	}
	dynamoClient := initializer.DynamoDbClient{
		Client:    dynamodb.NewFromConfig(cfg),
		TableName: "channeldata",
	}

	error, channel := db.ReadDataById(dynamoClient, c.Query("identifier"), c.Query("loggedInCity"))
	if error != nil {
		fmt.Printf("Error searching records")
		c.JSONP(http.StatusInternalServerError, error)
		return
	}
	c.JSONP(http.StatusOK, channel)

}
func deleteChannelById(c *gin.Context) {

	cfg, error := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))

	if error != nil {
		log.Fatalf("error acessing client %v", error)
	}
	dynamoClient := initializer.DynamoDbClient{
		Client:    dynamodb.NewFromConfig(cfg),
		TableName: "channeldata",
	}

	error = db.DeleteChannelById(dynamoClient, c.Query("identifier"), c.Query("loggedInCity"))
	if error != nil {
		fmt.Printf("Error deleting records")
		c.JSONP(http.StatusInternalServerError, error)
		return
	}
	c.JSONP(http.StatusOK, "Deleted successfully")

}
func addChannelAccess(c *gin.Context) {

	var customerChannelData model.CustomerChannel

	err := c.BindJSON(&customerChannelData)

	if err != nil {
		log.Fatal("Could not bind json")
	}

	cfg, error := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))

	if error != nil {
		log.Fatalf("error acessing client %v", error)
	}
	dynamoClient := initializer.DynamoDbClient{
		Client:    dynamodb.NewFromConfig(cfg),
		TableName: "channeldata",
	}
	dynamoChannel := model.Channel{
		Name:         customerChannelData.Name,
		Identifier:   customerChannelData.Identifier,
		CustomerId:   customerChannelData.CustomerId,
		LoggedInFrom: customerChannelData.LoggedInFrom,
	}

	err = db.StoreData(dynamoClient, dynamoChannel)
	if err != nil {
		fmt.Printf("Error integrating records")
		c.JSONP(http.StatusInternalServerError, err)
		return
	}
	c.JSONP(http.StatusOK, dynamoChannel)

}
