package app

import (
	"log"
	"os"
)

var (
	Setting Config
	Noti    *Api
)

const url = "https://notify-api.line.me/api/notify"

type Config struct {
	LineToken string `yaml:"LINE_TOKEN" env:"LINE_TOKEN"`
}

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
