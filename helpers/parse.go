package helpers

import (
	"encoding/json"
	"time"
)

func Parse[T any](data string) T {
	var result T
	_ = json.Unmarshal([]byte(data), &result)
	return result
}

func ParseTime(i string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z", i)
	return t
}