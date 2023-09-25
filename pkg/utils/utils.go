package utils

import (
	"log"
)

// ErrorHandler is a generic error handler
func ErrorHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}
