package main

type Message struct {
	ClientID string `json:"clientId"`
	Type     string `json:"type"`
	Data     string `json:"data,omitempty"`
	Time     int64  `json:"time"`
}
