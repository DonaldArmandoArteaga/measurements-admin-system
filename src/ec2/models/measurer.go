package models

import "time"

type Measurement struct {
	ID       string               `json:"id,omitempty"`
	Serial   string               `json:"serial,omitempty"`
	Date     time.Time            `json:"date,omitempty"`
	Values   interface{}          `json:"values,omitempty"`
	Metadata *MeasurementMetadata `json:"metadata,omitempty"`
}

type MeasurementMetadata struct {
	DateInserted time.Time `json:"DateInserted,omitempty"`
}

type MeasurementFromDynamo struct {
	ID       string `json:"id"`
	Serial   string `json:"serial"`
	Date     string `json:"date"`
	Values   string `json:"values"`
	Metadata string `json:"metadata"`
}
