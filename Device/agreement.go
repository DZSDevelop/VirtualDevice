package Device

import (
	. "VirtualDevice/client"
	"encoding/json"
)

//创建公共MAP
func CreateBaseMap(msgId string, command string) map[string]interface{} {
	base := make(map[string]interface{})
	if "" != msgId {
		base["msg_id"] = msgId
	}
	base["device_id"] = DeviceId
	base["auth_key"] = AuthKey
	base["command"] = command
	base["timestamp"] = GetTimestampMilli()
	return base
}

//心跳发送
func HeartBeat() ([]byte, error) {
	dataMap := CreateBaseMap("", "300000")
	return json.Marshal(dataMap)
}

//登录
func Login(cloudType string) ([]byte, error) {
	dataMap := CreateBaseMap("", "300002")
	data := make(map[string]interface{})
	data["cloud_type"] = cloudType
	dataMap["data"] = data
	return json.Marshal(dataMap)
}

//声波绑定用户
func SWBindUser(userId string, appId string) ([]byte, error) {
	dataMap := CreateBaseMap("", "300002")
	data := make(map[string]interface{})
	data["user_id"] = userId
	data["appid"] = appId
	dataMap["data"] = data
	return json.Marshal(dataMap)
}

//获取天气
func GetWeather(city string) ([]byte, error) {
	dataMap := CreateBaseMap("", "300022")
	data := make(map[string]interface{})
	data["city"] = city
	dataMap["data"] = data
	return json.Marshal(dataMap)
}
