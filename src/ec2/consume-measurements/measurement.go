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
)

type Measurement struct {
	Serial string      `json:"serial"`
	Date   time.Time   `json:"date"`
	Values interface{} `json:"values"`
}

func Init() {
	logs.InfoLogger.Println("Step 1")
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})
	logs.InfoLogger.Println("Step 2")

	if err != nil {
		logs.InfoLogger.Println("Step 3")
		logs.ErrorLogger.Println("Failed to initialize new session: ", err)
		return
	}

	logs.InfoLogger.Println("Step 4")
	sqsClient := sqs.New(sess)

	queueName := os.Getenv("QUEUE_NAME")

	urlRes, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	logs.InfoLogger.Println("Step 5")

	if err != nil {
		logs.InfoLogger.Println("Step 6")
		logs.ErrorLogger.Println("Got an error while trying to create queue: ", err)
		return
	}

	data := &Measurement{}

	logs.InfoLogger.Println("Step 7")
	svc := dynamodb.New(sess)

	for {
		logs.InfoLogger.Println("Step 8")
		time.Sleep(4 * time.Second)
		msgResult, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            urlRes.QueueUrl,
			MaxNumberOfMessages: aws.Int64(10),
		})
		logs.InfoLogger.Println("Step 9")

		if err != nil {
			logs.InfoLogger.Println("Step 10")
			logs.ErrorLogger.Println("Got an error while trying to retrieve message:", err)

			// TODO is that the behavior we want? what needs to happen when we get an error receiven a message?
			return
		}

		for _, message := range msgResult.Messages {
			logs.InfoLogger.Println("Step 11")
			go func(message *sqs.Message) {

				err := json.Unmarshal([]byte(*message.Body), data)
				if err != nil {
					logs.InfoLogger.Println("Step 12")
					logs.ErrorLogger.Println("Got an error while trying parse message into mesassuremnt struct: ", err)
					return
				}
				logs.InfoLogger.Println("Step 3")
				u, err := json.Marshal(data.Values)
				if err != nil {
					logs.InfoLogger.Println("Step 14")
					logs.ErrorLogger.Println("Got an error while trying parse message values into mesassuremnt values struct: ", err)
					return
				}

				_, err = svc.PutItem(&dynamodb.PutItemInput{
					TableName: aws.String(os.Getenv("DYNAMO_TABLE_NAME")),

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
				logs.InfoLogger.Println("Step 15")

				if err != nil {
					logs.InfoLogger.Println("Step 16")
					logs.ErrorLogger.Println("Got an error while trying to save message in dynamo: ", err)
					return
				}

				_, err = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      urlRes.QueueUrl,
					ReceiptHandle: message.ReceiptHandle,
				})
				logs.InfoLogger.Println("Step 17")

				if err != nil {
					logs.InfoLogger.Println("Step 18")
					logs.ErrorLogger.Println("Got an error while trying to delete message: ", err)
					return
				}

			}(message)

		}
	}

}
