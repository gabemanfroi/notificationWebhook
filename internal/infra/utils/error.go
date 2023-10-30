package utils

import (
	"log"
)

func HandleError(err error, message string) {
	if err != nil {
		log.Println(message + err.Error())
		panic(err)
	}
}
