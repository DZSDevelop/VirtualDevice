package main

import (
	"VirtualDevice/Device"
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
		var jsn []byte
		var err error
		switch scanner.Text() {
		case "0": //为0时结束程序
			return
		case "1": //登录
			jsn, err = Device.Login("2")
		case "2":
			jsn, err = Device.SWBindUser("628", "189c")
			break
		}
		sendMsg(c, jsn, err)
	}

}

//连接Server
func connectServer() (*Client, error) {
	c := NewClient(ClientId)
	err := c.Connect()
	return c, err
}

//发送消息
func sendMsg(c *Client, jb []byte, e error) {
	if e != nil {
		Println("Convert Err", e)
		return
	}
	Println("Send MSG:", string(jb))
	err := c.Publish(SendTopic, 0, false, jb)
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
	case "300022":
		handWeather(msg)
		break
	}
}

//处理登录逻辑
func handLogin(msg *Message) {

}

//处理天气信息
func handWeather(msg *Message) {

}
