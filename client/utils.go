package client

import (
	"fmt"
	"time"
)

func getTimestamp() int64 {
	return time.Now().Unix()
}
func getTimestampMilli() int64 {
	return getTimestampNano() / 1e6
}
func getTimestampNano() int64 {
	return time.Now().UnixNano()
}
func Println(a ...interface{}) {
	fmt.Println(a)
}
