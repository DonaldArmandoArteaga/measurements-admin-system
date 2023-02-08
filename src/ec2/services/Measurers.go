package services

import (
	"context"
	"encoding/json"
	"errors"
	"input-system/config"
	"input-system/models"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Measurers interface {
	GetById(id string) (*models.Measurement, error)
	GetBySerial(serial string) ([]*models.Measurement, int, error)
}

type MeasurersService struct {
	dynamoClient *dynamodb.Client
}

func CreateMeasurersService(dynamoClient *dynamodb.Client) *MeasurersService {
	return &MeasurersService{
		dynamoClient: dynamoClient,
	}
}

func (m *MeasurersService) GetById(id string) (*models.Measurement, error) {

	expr, err := expression.
		NewBuilder().
		WithKeyCondition(expression.Key("id").Equal(expression.Value(id))).
		Build()

	if err != nil {
		config.ErrorLogger.Println("Couldn't build expression for query: ", err)
		return nil, err
	}

	response, err := m.dynamoClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 &config.DYNAMO_TABLE_NAME,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		config.ErrorLogger.Println("Query has Failed: ", err)
		return nil, err
	}

	if len(response.Items) <= 0 {
		msg := "measurer not found"
		config.ErrorLogger.Println(msg)
		return nil, errors.New(msg)
	}

	var measurers []models.MeasurementFromDynamo
	err = attributevalue.UnmarshalListOfMaps(response.Items, &measurers)

	if err != nil {
		config.ErrorLogger.Println("Failed to unmarshal Record: ", err)
		return nil, err
	}

	date, err := time.Parse(config.TIME_FORMAT_1, measurers[0].Date)

	if err != nil {

		date, err = time.Parse(config.TIME_FORMAT_2, measurers[0].Date)

		if err != nil {
			config.ErrorLogger.Println("Failed to parte date from dynamo record: ", err)
			return nil, err
		}

	}

	metadata := &models.MeasurementMetadata{}
	err = json.Unmarshal([]byte(measurers[0].Metadata), &metadata)

	if err != nil {
		config.ErrorLogger.Println("Failed to parte metadata from dynamo record: ", err)
		return nil, err
	}

	return &models.Measurement{
		ID:       measurers[0].ID,
		Serial:   measurers[0].Serial,
		Date:     date,
		Values:   measurers[0].Values,
		Metadata: metadata,
	}, nil
}

func (m *MeasurersService) GetBySerial(serial string) ([]*models.Measurement, int, error) {

	expr, err := expression.
		NewBuilder().
		WithFilter(expression.Name("serial").Equal(expression.Value(serial))).
		Build()

	if err != nil {
		config.ErrorLogger.Println("Couldn't build expression for scan: ", err)
		return nil, 0, err
	}

	response, err := m.dynamoClient.Scan(context.Background(), &dynamodb.ScanInput{
		TableName:                 &config.DYNAMO_TABLE_NAME,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ConsistentRead:            aws.Bool(true),
	})

	if err != nil {
		config.ErrorLogger.Println("Query has Failed: ", err)
		return nil, 0, err
	}

	var measurersFromDynamo []models.MeasurementFromDynamo
	err = attributevalue.UnmarshalListOfMaps(response.Items, &measurersFromDynamo)

	if err != nil {
		config.ErrorLogger.Println("Failed to unmarshal Record: ", err)
		return nil, 0, err
	}

	var measurers []*models.Measurement
	var date time.Time

	for _, measurer := range measurersFromDynamo {

		date, err = time.Parse(config.TIME_FORMAT_1, measurer.Date)

		if err != nil {

			date, err = time.Parse(config.TIME_FORMAT_2, measurer.Date)

			if err != nil {
				config.ErrorLogger.Println("Failed to parte date from dynamo record: ", err)
				return nil, 0, err
			}

		}

		var measurerResponse *models.Measurement

		if measurer.Metadata != "" {

			metadata := &models.MeasurementMetadata{}
			err = json.Unmarshal([]byte(measurer.Metadata), &metadata)

			if err != nil {
				config.ErrorLogger.Println("Failed to parse metadata from dynamo record: ", err)
				return nil, 0, err
			}

			measurerResponse = &models.Measurement{
				ID:       measurer.ID,
				Serial:   measurer.Serial,
				Date:     date,
				Values:   measurer.Values,
				Metadata: metadata,
			}

		} else {

			measurerResponse = &models.Measurement{
				ID:     measurer.ID,
				Serial: measurer.Serial,
				Date:   date,
				Values: measurer.Values,
			}

		}

		measurers = append(measurers, measurerResponse)
	}

	return measurers, len(measurers), nil
}
