package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var (
	QUEUE_NAME        = os.Getenv("QUEUE_NAME")
	DYNAMO_TABLE_NAME = os.Getenv("DYNAMO_TABLE_NAME")
	AWS_REGION        = os.Getenv("AWS_REGION")
	TIME_FORMAT_1     = "2006-01-02 15:04:05.999999999 -0700 MST"
	TIME_FORMAT_2     = "2006-01-02 15:04:05.999999999 -0500 -0500"
)

func InitConfig() (*dynamodb.Client, *sqs.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(AWS_REGION),
	)

	if err != nil {
		ErrorLogger.Println("Failed to initialize new AWS session: ", err)
		return nil, nil, err
	}

	return dynamodb.NewFromConfig(cfg), sqs.NewFromConfig(cfg), nil
}
