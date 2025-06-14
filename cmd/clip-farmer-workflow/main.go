package main

import (
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/worker"
)

func main() {
	cfg := config.Config{
		TwitchHelixClientId:     config.GetEnv("HELIX_CLIENT_ID", ""),
		TwitchHelixSecret: config.GetEnv("HELIX_SECRET", ""),
		TwitchGQLClientId: config.GetEnv("GQL_CLIENT_ID", ""),
	}
	worker.StartWorker(cfg)
}
