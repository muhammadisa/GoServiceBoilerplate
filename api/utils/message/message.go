package message

import (
	"fmt"
	"strings"
)

// GenerateMessage generate error, successfully messages with status for consistency
func GenerateMessage(withID uint64, httpMethod string, targetModelName string, isSuccess bool) string {

	var with string
	if withID == 0 {
		with = "without id"
	} else {
		with = fmt.Sprintf("with id %d", withID)
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
				targetModelName,
				with,
			)
			break
		}
	}

	return messages

}
