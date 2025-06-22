package activity

import (
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/service/download"
	"clip-farmer-workflow/internal/service/helloworld"
	"clip-farmer-workflow/internal/service/twitch"
)

type Activity struct {
	helloworld.HelloWorldManager
	twitch.TwitchManager
	download.DownloadManager
}

func NewActivity(cfg config.Config) *Activity {
	twitchManager := twitch.NewTwitchService(cfg.TwitchHelixClientId, cfg.TwitchHelixSecret, cfg.TwitchGQLClientId)
	downloadManager := download.NewDownloadService()
	return &Activity{
		TwitchManager:     twitchManager,
		DownloadManager:  downloadManager,
	}
}
