package activity

import (
	"clip-farmer-workflow/internal/service/helloworld"
	"clip-farmer-workflow/internal/service/twitch"
)

type Activity struct {
	helloworld.HelloWorldManager
	twitch.TwitchManager
}