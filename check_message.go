package gateway_schema

import (
	"encoding/json"
	"errors"
	"fmt"
	valid "github.com/asaskevich/govalidator"
)

//验证message是否合法，并且返回序列化后的 message data
func CheckMessage(m BaseMessage) (b []byte, err error) {
	switch m.MessageType {

	case MessageTypePluginRegisterResponse:
		var data PluginRegisterResponse
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	case MessageTypeDeviceSavedNotification:
		var data DeviceSavedNotification
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	case MessageTypeAdapterRemoveDeviceRequest:
		var data AdapterRemoveDeviceRequest
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	case MessageTypeAdapterRemoveDeviceResponse:
		var data AdapterRemoveDeviceResponse
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	case MessageTypeDeviceSetPropertyCommand:
		var data DeviceSetPropertyCommand
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	case MessageTypePluginRegisterRequest:
		var data PluginRegisterRequest
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	case MessageTypePluginUnloadRequest:
		var data PluginUnloadRequest
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	case MessageTypePluginErrorNotification:
		var data PluginErrorNotification
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	case MessageTypeDevicePropertyChangedNotification:
		var data DevicePropertyChangedNotification
		b, _ = json.Marshal(m.Data)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return b, err
		}
		_, err = valid.ValidateStruct(data)
	default:
		b = nil
		err = errors.New(fmt.Sprintf("message unknown: %v", m))
	}
	return b, err
}
