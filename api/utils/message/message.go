package message

import (
	"fmt"
	"reflect"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// GetType get struct name
func GetType(myvar interface{}) string {
	valueOf := reflect.ValueOf(myvar)
	if valueOf.Type().Kind() == reflect.Ptr {
		return reflect.Indirect(valueOf).Type().Name()
	}
	return valueOf.Type().Name()
}

// GenerateMessage generate error, successfully messages with status for consistency
func GenerateMessage(
	withID uuid.UUID,
	httpMethod string,
	model interface{},
	isSuccess bool,
) string {

	var with string
	if withID == uuid.Nil {
		with = "without id"
	} else {
		with = fmt.Sprintf("with id %s", withID.String())
	}

	var kind string
	if isSuccess {
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
		if httpMethod == mtd[0] {
			messages = fmt.Sprintf(
				"%s %s %s data %s",
				kind,
				mtd[1],
				GetType(model),
				with,
			)
			break
		}
	}

	return messages

}
