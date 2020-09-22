package gateway_schema

const (
	MessageTypePluginRegisterRequest   = 0 //注册
	MessageTypePluginRegisterResponse  = 1
	MessageTypePluginUnloadRequest     = 2 //卸载
	MessageTypePluginUnloadResponse    = 3
	MessageTypePluginErrorNotification = 4
)

// plugin register
type PluginRegisterRequest struct {
	PluginId string `validate:"required"  json:"plugin_id"`
}

// plugin register response
type PluginRegisterResponse struct {
	PluginId    string
	UserProfile *UserProfile `json:"user_profile"`
	Preferences *Preferences `json:"preferences"`
}

type PluginUnloadRequest struct {
	PluginId string `validate:"required"  json:"plugin_id"`
}

type PluginUnloadResponse struct {
	PluginId string `validate:"required"  json:"plugin_id"`
}

type PluginErrorNotification struct {
	PluginId string `validate:"required"  json:"plugin_id"`
	Message  string `json:"message"`
}
