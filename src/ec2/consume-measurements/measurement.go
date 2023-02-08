package consumemeasurements

import (
	"context"
	"encoding/json"
	"input-system/config"
	"input-system/models"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"

	"github.com/google/uuid"
)

func Init(svc *dynamodb.Client, sqsClient *sqs.Client) {

	urlRes, err := sqsClient.GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
		QueueName: &config.QUEUE_NAME,
	})

	if err != nil {
		config.ErrorLogger.Println("Got an error while trying to create queue: ", err)
		return
	}

	data := &models.Measurement{}

	config.InfoLogger.Println("Starting the infinitive loop...")

	for {
		time.Sleep(4 * time.Second)

		msgResult, err := sqsClient.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
			QueueUrl:            urlRes.QueueUrl,
			MaxNumberOfMessages: int32(10),
			MessageAttributeNames: []string{
				string(types.QueueAttributeNameAll),
			},
		})

		if err != nil {
			config.ErrorLogger.Println("Got an error while trying to retrieve message:", err)
			// TODO is that the behavior we want? what needs to happen when we get an error receiven a message?
			return
		}

		for _, message := range msgResult.Messages {
			go func(message *types.Message) {

				config.InfoLogger.Println("message:", *message.Body)
				err := json.Unmarshal([]byte(*message.Body), data)
				if err != nil {
					config.ErrorLogger.Println("Got an error while trying parse message into mesassuremnt struct: ", err)
					return
				}

				data.ID = uuid.New().String()
				data.Metadata = &models.MeasurementMetadata{DateInserted: time.Now()}
				item, err := attributevalue.MarshalMap(data)

				if err != nil {
					config.ErrorLogger.Println("Got an error while trying to convert the message into attribute value: ", err)
					return
				}

				_, err = svc.PutItem(context.Background(), &dynamodb.PutItemInput{
					TableName: &config.DYNAMO_TABLE_NAME,
					Item:      item,
				})

				if err != nil {
					config.ErrorLogger.Println("Got an error while trying to save message in dynamo: ", err)
					return
				}

				_, err = sqsClient.DeleteMessage(context.Background(), &sqs.DeleteMessageInput{
					QueueUrl:      urlRes.QueueUrl,
					ReceiptHandle: message.ReceiptHandle,
				})

				if err != nil {
					config.ErrorLogger.Println("Got an error while trying to delete message: ", err)
					return
				}

			}(&message)

		}
	}

}
