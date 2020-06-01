package main

import (
	. "VirtualDevice/client"
)

const (
	DeviceId  = "2000146200013882"
	SendTopic = "data_from_server/:" + DeviceId
	SubTopic  = "data_from_client"
	ClientId  = "DEV:" + DeviceId
)

func main() {
	var c = NewClient(ClientId)
	err := c.Connect()
	if err != nil {
		print("链接失败", err)
		return
	}
	_ = c.Subscribe(func(client *Client, msg *Message) {
		print(msg.Data)
	}, 0, SubTopic)
	js := "{\"msg_id\":\"\",\"command\":\"300002\",\"device_id\":\"2000146200013882\",\"timestamp\":\"1551148989\",\"data\":{\"cloud_type\":\"1\"}}"
	c.Publish(SendTopic, 0, false, []byte(js))
}
