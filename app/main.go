package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/setting"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/usecase"
)

func main() {
	ctx := context.Background()
	us, err := usecase.InitSendNotification()
	if err != nil {
		log.Fatal(err)
	}

	schedule := os.Getenv("SCHEDULE")
	if schedule == "" {
		log.Fatal(fmt.Errorf("notification schedule is not setted"))
	}
	fmt.Printf("exec schedule: %s\n", schedule)

	ns := filterNotifications(setting.Get().Notifications, schedule)
	for _, n := range ns {
		err = us.Do(ctx, n)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func filterNotifications(ns []setting.Notification, schedule string) []setting.Notification {
	fns := []setting.Notification{}
	for _, n := range ns {
		if n.Schedule == schedule {
			fns = append(fns, n)
		}
	}
	return fns
}
