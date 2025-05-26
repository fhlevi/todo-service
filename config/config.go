package config

import "log"

var AppPort = ":8000"

func Init() {
	log.Println("Configuration loaded")
}