package activity

import (
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/service/download"
	"clip-farmer-workflow/internal/service/edit"
	"clip-farmer-workflow/internal/service/helloworld"
	"clip-farmer-workflow/internal/service/twitch"
)

type Activity struct {
	helloworld.HelloWorldManager
	twitch.TwitchManager
	download.DownloadManager
	edit.EditManager
}

func NewActivity(cfg config.Config) *Activity {
	helloManager := helloworld.HelloWorldService{}
	twitchManager := twitch.NewTwitchService(cfg.TwitchBaseURL, cfg.TwitchClientId, cfg.TwitchClientSecret)
	downloadManager := download.NewDownloadService()
	editManager := edit.EditService{}
	return &Activity{
		HelloWorldManager: helloManager,
		TwitchManager:     twitchManager,
		DownloadManager:  downloadManager,
		EditManager: editManager,
	}
}
