package checkdebug

import (
	"os"
	"strconv"
)

// CheckDebug checking debug key environment
func CheckDebug() bool {
	// Load debuging mode env
	debugEnv := os.Getenv("DEBUG")
	debug, err := strconv.ParseBool(debugEnv)
	if err != nil {
		return false
	}
	return debug
}
