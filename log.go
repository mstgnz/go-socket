package main

import (
	"log"
)

func HandleLog(message string, err error) {
	if err != nil {
		log.Println(message, err.Error())
	} else {
		log.Println(message)
	}
}
