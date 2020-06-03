package client

import (
	"fmt"
	"time"
)

func GetTimestamp() int64 {
	return time.Now().Unix()
}
func GetTimestampMilli() int64 {
	return GetTimestampNano() / 1e6
}
func GetTimestampNano() int64 {
	return time.Now().UnixNano()
}
func Println(a ...interface{}) {
	fmt.Println(a)
}
