package commons

import (
	 "time"
)

const (
	ApplicationJsonContentType = "application/json; charset=utf-8"
)

func InitializeString(value string) *string {
	return &value
}

func InitializeBool(value bool) *bool {
	return &value
}

func InitializeTime(value time.Time) *time.Time {
	return &value
}