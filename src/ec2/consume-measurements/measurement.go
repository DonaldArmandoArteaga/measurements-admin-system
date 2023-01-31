package consumemeasurements

import (
	"encoding/json"
	logs "input-system/config"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

type Measurement struct {
	Serial string      `json:"serial"`
	Date   time.Time   `json:"date"`
	Values interface{} `json:"values"`
}

func Init() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		logs.ErrorLogger.Println("Failed to initialize new session: ", err)
		return
	}

	sqsClient := sqs.New(sess)

	queueName := os.Getenv("QUEUE_NAME")

	urlRes, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		logs.ErrorLogger.Println("Got an error while trying to create queue: ", err)
		return
	}

	data := &Measurement{}

	svc := dynamodb.New(sess)

	for {
		time.Sleep(4 * time.Second)
		msgResult, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            urlRes.QueueUrl,
			MaxNumberOfMessages: aws.Int64(10),
		})

		if err != nil {
			logs.ErrorLogger.Println("Got an error while trying to retrieve message:", err)
			// TODO is that the behavior we want? what needs to happen when we get an error receiven a message?
			return
		}

		for _, message := range msgResult.Messages {
			go func(message *sqs.Message) {

				logs.InfoLogger.Println("message:", *message.Body)
				err := json.Unmarshal([]byte(*message.Body), data)
				if err != nil {
					logs.ErrorLogger.Println("Got an error while trying parse message into mesassuremnt struct: ", err)
					return
				}
				u, err := json.Marshal(data.Values)
				if err != nil {
					logs.ErrorLogger.Println("Got an error while trying parse message values into mesassuremnt values struct: ", err)
					return
				}

				_, err = svc.PutItem(&dynamodb.PutItemInput{
					TableName: aws.String(os.Getenv("DYNAMO_TABLE_NAME")),

					Item: map[string]*dynamodb.AttributeValue{
						"id": {
							S: aws.String(uuid.New().String()),
						},
						"serial": {
							S: aws.String(data.Serial),
						},
						"date": {
							S: aws.String(data.Date.String()),
						},
						"values": {
							S: aws.String(string(u)),
						},
					},
				})

				if err != nil {
					logs.ErrorLogger.Println("Got an error while trying to save message in dynamo: ", err)
					return
				}

				_, err = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      urlRes.QueueUrl,
					ReceiptHandle: message.ReceiptHandle,
				})

				if err != nil {
					logs.ErrorLogger.Println("Got an error while trying to delete message: ", err)
					return
				}

			}(message)

		}
	}

}
