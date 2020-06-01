package client

type Message struct {
	MsgId     string                 `json:"msg_id,omitempty"`
	Command   string                 `json:"command"`
	DeviceId  string                 `json:"device_id"`
	Timestamp string                 `json:"timestamp"`
	Data      map[string]interface{} `json:"data,omitempty"`
}
