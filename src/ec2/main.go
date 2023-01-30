package main

import (
	"fmt"
	consumemeasurements "input-system/consume-measurements"
	"os"
)

func main() {
	consumemeasurements.Init()
	fmt.Println("QueueName:", os.Getenv("QUEUE_NAME"))
	fmt.Println("TableName:", os.Getenv("DYNAMO_TABLE_NAME"))
	fmt.Println("Finish Init Main")
}
