package message

import "fmt"

// IMessage interface
type IMessage interface {
	GenerateMessage() string
}

// Message struct
type Message struct {
	WithID          int64
	HTTPMethod      string
	IsSuccess       bool
	TargetModelName string
}

// GenerateMessage generate error, successfully messages with status for consistency
func (msg Message) GenerateMessage() string {

	var with string
	if msg.WithID == 0 {
		with = "without id"
	} else {
		with = fmt.Sprintf("with id %d", msg.WithID)
	}

	var kind string
	if msg.IsSuccess {
		kind = "Successfully"
	} else {
		kind = "Failed to"
	}

	var message string

	switch msg.HTTPMethod {
	case "GET":
		message = fmt.Sprintf(
			"%s retrieved %s data %s",
			kind,
			msg.TargetModelName,
			with,
		)
	case "POST":
		message = fmt.Sprintf(
			"%s saved %s data %s",
			kind,
			msg.TargetModelName,
			with,
		)
	case "PUT":
		message = fmt.Sprintf(
			"%s edit %s data %s",
			kind,
			msg.TargetModelName,
			with,
		)
	case "PATCH":
		message = fmt.Sprintf(
			"%s edit %s data %s",
			kind,
			msg.TargetModelName,
			with,
		)
	case "DELETE":
		message = fmt.Sprintf(
			"%s deleted %s data %s",
			kind,
			msg.TargetModelName,
			with,
		)
	default:
		message = ""
	}

	return message
}
