package message

import (
	"fmt"
	"strings"
)

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

	methods := []string{
		"GET,retrieved",
		"POST,saved",
		"PUT,edit",
		"PATCH,edit",
		"DELETE,deleted",
	}

	var messages string

	for index := range methods {
		mtd := strings.Split(methods[index], ",")
		if msg.HTTPMethod == mtd[0] {
			messages = fmt.Sprintf(
				"%s %s %s data %s",
				kind,
				mtd[1],
				msg.TargetModelName,
				with,
			)
			break
		}
	}

	return messages

}
