package main

import (
	"input-system/config"
	logs "input-system/config"
	consumemeasurements "input-system/consume-measurements"
	"input-system/controllers"
	"input-system/services"

	"github.com/gin-gonic/gin"
)

func main() {
	logs.InfoLogger.Println("Starting the application...")
	logs.InfoLogger.Println("QueueName:", config.QUEUE_NAME)
	logs.InfoLogger.Println("DynamoTableName:", config.DYNAMO_TABLE_NAME)
	logs.InfoLogger.Println("AWSRegion:", config.AWS_REGION)

	dynamoClient, sqsClient, err := config.InitConfig()

	if err != nil {
		panic(err)
	}

	measurerService := services.CreateMeasurersService(dynamoClient)
	go consumemeasurements.Init(dynamoClient, sqsClient)

	r := gin.Default()
	controllers.CreateHealthCheckController(r)
	controllers.CreateMeasurersController(r, measurerService)
	r.Run()

}
