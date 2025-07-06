package main

import (
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/worker"
)

func main() {
	cfg := config.Config{
		TwitchGQLClientId: config.GetEnv("TWITCH_GQL_CLIENT_ID", ""),
	}
	worker.StartWorker(cfg)
}
