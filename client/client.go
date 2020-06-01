package client

import (
	"encoding/json"
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"sync"
	"time"
)

const (
	HOST          = "wmq.worthcloud.net:1883"
	PORT          = 1883
	USERNAME      = ""
	PWD           = ""
	KEEPALIVE     = 120 * time.Second
	PING_TIMEOUT  = 10 * time.Second
	WRITE_TIMEOUT = 10 * time.Second
	WAIT_TIMEOUT  = 10 * time.Second
)

type Client struct {
	client        mqtt.Client
	clientOptions *mqtt.ClientOptions
	mutex         *sync.Mutex
	observer      func(c *Client, msg *Message)
}

func NewClient(clientId string) *Client {
	clientOptions := mqtt.NewClientOptions().
		AddBroker(HOST).
		//SetUsername(USERNAME).
		//SetPassword(PWD).
		SetClientID(clientId).
		SetCleanSession(false).
		SetAutoReconnect(true).
		SetKeepAlive(KEEPALIVE).
		SetPingTimeout(PING_TIMEOUT).
		SetWriteTimeout(WRITE_TIMEOUT).
		SetOnConnectHandler(func(client mqtt.Client) {
			Println("MQTT is connected,ClientId is", clientId)
		}).
		SetConnectionLostHandler(func(client mqtt.Client, err error) {
			Println("MQTT is disconnected,ClientId is", clientId)
			Println("Error is", err)
		})
	client := mqtt.NewClient(clientOptions)
	return &Client{
		client:        client,
		clientOptions: clientOptions,
		mutex:         &sync.Mutex{},
	}
}
func (c *Client) GetClientId() string {
	return c.clientOptions.ClientID
}
func (c *Client) Connect() error {
	if !c.client.IsConnected() {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		if !c.client.IsConnected() {
			if token := c.client.Connect(); token.Wait() && token.Error() != nil {
				return token.Error()
			}
		}
	}
	return nil
}
func (c *Client) Publish(topic string, qos byte, retained bool, data []byte) error {
	if err := c.Connect(); err != nil {
		return err
	}
	token := c.client.Publish(topic, qos, retained, data)
	if err := token.Error(); err != nil {
		return err
	}
	if !token.WaitTimeout(WAIT_TIMEOUT) {
		return errors.New("MQTT publish wait timeout")
	}
	return nil
}

//订阅消息
func (c *Client) Subscribe(observer func(c *Client, msg *Message), qos byte, topics ...string) error {
	if len(topics) == 0 {
		return errors.New("the topic is empty")
	}
	if observer == nil {
		return errors.New("the observer func is nil")
	}
	if c.observer != nil {
		return errors.New("an existing observer subscribed on this client, you must unsubscribe it before you subscribe a new observer")
	}
	c.observer = observer
	filters := make(map[string]byte)
	for _, topic := range topics {
		filters[topic] = qos
	}
	c.client.SubscribeMultiple(filters, c.messageHandler)
	return nil
}

//消息处理
func (c *Client) messageHandler(client mqtt.Client, msg mqtt.Message) {
	if c.observer == nil {
		Println("not subscribe message observer")
		return
	}
	Println("message received", string(msg.Payload()))
	message, err := decodeMessage(msg.Payload())
	if err != nil {
		Println("failed to decode message")
		return
	}
	c.observer(c, message)
}

//解析消息
func decodeMessage(payload []byte) (*Message, error) {
	message := new(Message)
	decoder := json.NewDecoder(strings.NewReader(string(payload)))
	decoder.UseNumber()
	if err := decoder.Decode(&message); err != nil {
		return nil, err
	}
	return message, nil
}

//取消订阅
func (c *Client) Unsubscribe(topics ...string) {
	c.observer = nil
	c.client.Unsubscribe(topics...)
}
