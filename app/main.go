package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"

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

	if len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("notification not selected"))
	}

	ntfName := os.Args[1]
	sntf, err := filterNotification(setting.Get().Notifications, ntfName)
	if err != nil {
		log.Fatal(err)
	}

	err = us.Do(ctx, sntf)
	if err != nil {
		log.Fatal(err)
	}
}

func filterNotification(ntfs []setting.Notification, name string) (setting.Notification, error) {
	for _, ntf := range ntfs {
		if ntf.Name == name {
			return ntf, nil
		}
	}
	return setting.Notification{}, fmt.Errorf("notification: %s is not found", name)
}
