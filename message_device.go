package gateway_schema

const (
	MessageTypeAdapterRemoveDeviceRequest  = 4103
	MessageTypeAdapterRemoveDeviceResponse = 4104

	MessageTypeDeviceSavedNotification = 8207
	MessageTypeDeviceAddedNotification = 8192

	MessageTypeDeviceSetPinRequest  = 8193
	MessageTypeDeviceSetPinResponse = 8194

	MessageTypeDeviceSetPropertyCommand          = 8198
	MessageTypeDevicePropertyChangedNotification = 8199

	MessageTypeDeviceSetCredentialsRequest  = 8195
	MessageTypeDeviceSetCredentialsResponse = 8196
)

type DeviceSavedNotification struct {
	PluginId  string      `json:"plugin_id"`
	AdapterId string      `json:"adapter_id"`
	DeviceId  string      `json:"device_id"`
	Device    interface{} `json:"device"`
}

type DeviceSetPropertyCommand struct {
	PluginId      string      `json:"plugin_id"`
	AdapterId     string      `json:"adapter_id"`
	DeviceId      string      `json:"device_id"`
	PropertyName  string      `json:"property_name"`
	PropertyValue interface{} `json:"property_value"`
}

type DeviceAddedNotification struct {
	PluginId  string      `json:"plugin_id"`
	AdapterId string      `json:"adapter_id"`
	Device    interface{} `json:"device"`
}

type DevicePropertyChangedNotification struct {
	AdapterId string `json:"adapter_id"`
	DeviceId  string `json:"device_id"`
	Property  string `json:"property"`
}

type DeviceSetPinRequest struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
	MessageId int    `json:"message_id"`
	DeviceId  string `json:"device_id"`
	Pin       string `json:"pin"`
}

type DeviceSetPinResponse struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
	MessageId int    `json:"message_id"`
	DeviceId  string `json:"device_id"`
	Device    string `json:"device"`
	Success   bool   `json:"success"`
}

type DeviceSetCredentialsRequest struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
	MessageId int    `json:"message_id"`
	DeviceId  string `json:"device_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type DeviceSetCredentialsResponse struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
	DeviceId  string `json:"device_id"`
	Device    string `json:"device"`
	MessageId int    `json:"message_id"`
	Success   bool   `json:"success"`
}
