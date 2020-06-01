package main

import (
	. "VirtualDevice/client"
	"sync"
)

const (
	DeviceId  = "2000146200013882"
	SendTopic = "data_from_client"
	SubTopic  = "data_from_server/" + DeviceId
	ClientId  = "DEV:" + DeviceId
)

func main() {
	var wg = sync.WaitGroup{}
	var c = NewClient(ClientId)
	err := c.Connect()
	if err != nil {
		Println("链接失败", err)
		return
	}
	wg.Add(1)
	_ = c.Subscribe(func(client *Client, msg *Message) {
		Println(msg.Data)
	}, 0, SubTopic)

	js := "{\"msg_id\":\"\",\"command\":\"300002\",\"device_id\":\"2000146200013882\",\"timestamp\":\"1551148989\",\"data\":{\"cloud_type\":\"1\"}}"
	err = c.Publish(SendTopic, 0, false, []byte(js))
	if err != nil {
		Println(err)
	}
	wg.Wait()

}
