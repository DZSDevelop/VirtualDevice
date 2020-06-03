package Device

import (
	. "VirtualDevice/client"
	"encoding/json"
)

func CreateBaseMap(msgId string, command string) map[string]interface{} {
	base := make(map[string]interface{})
	if "" != msgId {
		base["msg_id"] = msgId
	}
	base["device_id"] = DeviceId
	base["timestamp"] = GetTimestampMilli()
	return base
}

func HeartBeat() ([]byte, error) {
	dataMap := CreateBaseMap("", "30000")
	return json.Marshal(dataMap)
}
func Login(cloudType string) ([]byte, error) {
	dataMap := CreateBaseMap("", "30002")
	dataMap["cloud_type"] = cloudType
	return json.Marshal(dataMap)
}
func BindUser() {

}
