package activity

import "clip-farmer-workflow/internal/service/twitch"

type Activity struct {
	twitchService twitch.TwitchManager
}