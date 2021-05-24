package errors

import "log"

// HandleErr panics and logs the error
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
