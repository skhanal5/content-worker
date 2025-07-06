package activity

import (
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/service/download"
	"clip-farmer-workflow/internal/service/twitch"
)

type Activity struct {
	twitch.TwitchManager
	download.DownloadManager
}

func NewActivity(cfg config.Config) *Activity {
	twitchManager := twitch.NewTwitchService(cfg.TwitchGQLClientId)
	downloadManager := download.NewDownloadService()
	return &Activity{
		TwitchManager:   twitchManager,
		DownloadManager: downloadManager,
	}
}
