package main

import (
	. "VirtualDevice/client"
	"bufio"
	"os"
)

func main() {
	c, err := connectServer()
	if err != nil {
		Println("链接失败", err)
		return
	}
	_ = c.Subscribe(handMsg, 0, SubTopic)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "0": //为0时结束程序
			return
		case "1":
		case "2":
			break
		}
	}

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

//处理消息
func handMsg(client *Client, msg *Message) {
	switch msg.Command {
	//登录返回
	case "300002":
		handLogin(msg)
		break
	case "300013":
		break
	}
}
func handLogin(msg *Message) {

}
