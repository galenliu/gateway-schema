package gateway_schema

const (
	MessageTypeAdapterRemoveDeviceRequest  = 4103
	MessageTypeAdapterRemoveDeviceResponse = 4104

	MessageTypeDeviceSavedNotification = 8207
	MessageTypeDeviceAddedNotification = 8192

	MessageTypeDeviceSetPropertyCommand          = 8198
	MessageTypeDevicePropertyChangedNotification = 8199
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
