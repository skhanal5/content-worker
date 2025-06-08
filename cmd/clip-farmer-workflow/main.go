package main

import (
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/worker"
)

func main() {
	cfg := config.Config{
		TwitchClientId:     config.GetEnv("TWITCH_CLIENT_ID", ""),
		TwitchClientSecret: config.GetEnv("TWITCH_CLIENT_SECRET", ""),
	}
	worker.StartWorker(cfg)
}
