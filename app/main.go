package main

import (
	"context"
	"embed"
	"log"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/setting"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/usecase"
)

// go:embed queries/*
var QueryFiles embed.FS

// go:embed setting/*
var SettingFiles embed.FS

func main() {
	ctx := context.Background()
	us, err := usecase.InitSendNotification()
	if err != nil {
		log.Fatal(err)
	}
	sntf := setting.Get().Notifications[0]
	err = us.Do(ctx, sntf)
	if err != nil {
		log.Fatal(err)
	}
}
