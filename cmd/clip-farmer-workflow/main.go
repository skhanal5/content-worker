package main

import (
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/worker"
)

func main() {
	cfg := config.Config{
        TwitchClientId:    config.GetEnv("TWITCH_CLIENT_ID", ""),
		TwitchClientSecret: config.GetEnv("TWITCH_CLIENT", ""),
		TwitchBaseURL:     config.GetEnv("TWITCH_BASE_URL", "https://api.twitch.tv/helix"),
    }
	worker.StartWorker(cfg)
}