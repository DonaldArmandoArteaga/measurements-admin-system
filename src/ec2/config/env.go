package config

import "os"

var (
	QUEUE_NAME        = os.Getenv("QUEUE_NAME")
	DYNAMO_TABLE_NAME = os.Getenv("DYNAMO_TABLE_NAME")
)
