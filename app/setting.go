package app

import (
	"log"
	"os"
)

var (
	Setting Config
	Noti    *Api

	// status code channel
	TrackingChannel chan int
	// bocy message channel
	MessageChannel chan string
)

const url = "https://notify-api.line.me/api/notify"

type Config struct {
	LineToken string `yaml:"LINE_TOKEN" env:"LINE_TOKEN"`
}

// InitSetting set variable and initialize struct
func InitSetting() {
	var err error

	Setting.LineToken = os.Getenv("LINE_TOKEN")
	if err != nil {
		log.Fatalf("could not get environment variable. %v", err)
	}

	Noti = &Api{
		Url:   url,
		Token: Setting.LineToken,
	}
}
