package gateway_schema

const (
	MessageTypeAdepterAddedNotification = 4096

	MessageTypeAdapterCancelRemoveDeviceCommand = 4105

	MessageTypeAdapterStartPairingCommand  = 4099
	MessageTypeAdapterCancelPairingCommand = 4100
	MessageTypeAdapterUnloadRequest        = 4097
	MessageTypeAdapterUnloadResponse       = 4098
)

type AdapterAddedNotification struct {
	PackageName string `validate:"required" json:"package_name"`
	AdapterId   string `validate:"required" json:"adapter_id"`
	Name        string `validate:"required" json:"name"`
}

type AdapterStartPairingCommand struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
	Timeout   int    `json:"timeout"`
}

type AdapterCancelPairingCommand struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
}

type AdapterUnloadRequest struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
}

type AdapterUnloadResponse struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
}

type AdapterRemoveDeviceRequest struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
	DeviceId  string `json:"device_id"`
}

type AdapterRemoveDeviceResponse struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
	DeviceId  string `json:"device_id"`
}

type AdapterCancelRemoveDeviceCommand struct {
	PluginId  string `json:"plugin_id"`
	AdapterId string `json:"adapter_id"`
	DeviceId  string `json:"device_id"`
}
