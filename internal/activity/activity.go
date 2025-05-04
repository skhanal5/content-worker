package activity

import (
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/service/helloworld"
	"clip-farmer-workflow/internal/service/twitch"
)

type Activity struct {
	helloworld.HelloWorldManager
	twitch.TwitchManager
}

func NewActivity(cfg config.Config) *Activity {
	twitchManager := twitch.NewTwitchService(cfg.TwitchBaseURL, cfg.TwitchClientId, cfg.TwitchClientSecret)
	return &Activity{
		HelloWorldManager: helloworld.HelloWorldService{},
		TwitchManager:     twitchManager,
	}
}
