package handlers

import "log"

func HandleError(err error) bool {
	if err != nil {
		log.Printf("[ERROR] %v", err.Error())
		return true
	}
	return false
}
