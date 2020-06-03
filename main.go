package main

import (
	. "VirtualDevice/client"
	"sync"
)

func main() {
	c, err := connectServer()
	if err != nil {
		Println("链接失败", err)
		return
	}
	var wg = sync.WaitGroup{}
	wg.Add(1)
	_ = c.Subscribe(func(client *Client, msg *Message) {
		Println(msg.Data)

	}, 0, SubTopic)
	wg.Wait()
}

//连接Server
func connectServer() (*Client, error) {
	c := NewClient(ClientId)
	err := c.Connect()
	return c, err
}

//发送消息
func sendMsg(c *Client, json string) {
	err := c.Publish(SendTopic, 0, false, []byte(json))
	if err != nil {
		Println("Send MSG Error:", err)
	}
}
