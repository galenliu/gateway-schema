package gateway_schema

type BaseMessage struct {
	MessageType int         `json:"message_type" validate:"required"`
	Data        interface{} `json:"data" validate:"required"`
}

//// Plugin Adapter Added Notification
