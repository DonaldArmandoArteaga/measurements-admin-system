package main

import (
	"input-system/config"
	logs "input-system/config"
	consumemeasurements "input-system/consume-measurements"
)

func main() {
	logs.InfoLogger.Println("Starting the application...")
	logs.InfoLogger.Println("QueueName:", config.QUEUE_NAME)
	logs.InfoLogger.Println("DynamoTableName:", config.DYNAMO_TABLE_NAME)
	logs.InfoLogger.Println("AWSRegion:", config.AWS_REGION)
	consumemeasurements.Init()
}
