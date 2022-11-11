package logRequest

type LogSendAction struct {
	RabbitMQ map[string]string `json:"rabbitMQ"`
	Response map[string]string `json:"response"`
}
