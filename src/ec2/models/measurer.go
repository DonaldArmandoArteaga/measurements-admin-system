package models

import "time"

type Measurement struct {
	Id       string               `json:"id,omitempty" dynamodbav:"id"`
	Serial   string               `json:"serial,omitempty" dynamodbav:"serial"`
	Date     time.Time            `json:"date,omitempty" dynamodbav:"date"`
	Values   interface{}          `json:"values,omitempty" dynamodbav:"values"`
	Metadata *MeasurementMetadata `json:"metadata,omitempty" dynamodbav:"metadata"`
}

type MeasurementMetadata struct {
	DateInserted time.Time `json:"DateInserted,omitempty" dynamodbav:"DateInserted"`
}

type MeasurementFromDynamo struct {
	ID       string `json:"id"`
	Serial   string `json:"serial"`
	Date     string `json:"date"`
	Values   string `json:"values"`
	Metadata string `json:"metadata"`
}
