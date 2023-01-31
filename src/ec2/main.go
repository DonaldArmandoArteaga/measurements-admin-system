package main

import (
	logs "input-system/config"
	consumemeasurements "input-system/consume-measurements"

	"os"
)

func main() {
	consumemeasurements.Init()
	logs.InfoLogger.Println("Starting the application...")
	logs.InfoLogger.Println("QueueName:", os.Getenv("QUEUE_NAME"))
	logs.InfoLogger.Println("TableName:", os.Getenv("DYNAMO_TABLE_NAME"))
	logs.InfoLogger.Println("Finish Init Main")

}
