package consumemeasurements

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
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
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	sqsClient := sqs.New(sess)

	queueName := os.Getenv("QUEUE_NAME") //"InputSystemStack-InputSystemQueueD5E56904-yDmemu8H2zaN"

	urlRes, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		fmt.Printf("Got an error while trying to create queue: %v", err)
		return
	}

	data := &Measurement{}

	svc := dynamodb.New(sess)

	//a := "https://sqs.us-east-1.amazonaws.com/478317648480/InputSystemStack-InputSystemQueueD5E56904-yDmemu8H2zaN"
	for {
		time.Sleep(4 * time.Second)
		msgResult, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            urlRes.QueueUrl,
			MaxNumberOfMessages: aws.Int64(10),
		})

		if err != nil {
			fmt.Printf("Got an error while trying to retrieve message: %v", err)

			// TODO is that the behavior we want? what needs to happen when we get an error receiven a message?
			return
		}

		for _, message := range msgResult.Messages {

			go func(message *sqs.Message) {

				err := json.Unmarshal([]byte(*message.Body), data)
				if err != nil {
					fmt.Printf("Got an error while trying parse message into mesassuremnt struct: %v  \n", err)
					return
				}

				u, err := json.Marshal(data.Values)
				if err != nil {
					fmt.Printf("Got an error while trying parse message values into mesassuremnt values struct: %v  \n", err)
					return
				}

				_, err = svc.PutItem(&dynamodb.PutItemInput{
					TableName: aws.String(os.Getenv("DYNAMO_TABLE_NAME")), //aws.String("InputSystemStack-InputSystemDynamoTable81679F4B-1WD4RD6LUN9YU"),
					Item: map[string]*dynamodb.AttributeValue{
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
					fmt.Printf("Got an error while trying to save message in dynamo: %v  \n", err)
					return
				}

				_, err = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      urlRes.QueueUrl,
					ReceiptHandle: message.ReceiptHandle,
				})

				if err != nil {
					fmt.Printf("Got an error while trying to delete message: %v  \n", err)
					return
				}

			}(message)

		}
	}

}
