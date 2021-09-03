package handlers

import "log"

func PanicOnError(err error) {
	if err == nil {
		return
	}
	log.Panicln(err)
}
