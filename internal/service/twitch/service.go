package twitch

import (
	"time"
	"github.com/go-resty/resty/v2"
)

type TwitchManager interface {
	GetUsers(broadcasterId string) (*UsersResponse, error)
	GetClips(username string, daysAgo int) (*ClipsResponse, error)
	GetDownloadLink(clipSlug string) (string, error)
}

const (
    baseUrl = "https://api.twitch.tv/helix"
	gqlUrl = "https://gql.twitch.tv"
	oauthBaseURL = "https://id.twitch.tv/oauth2/token"
)


type TwitchService struct {
	helixClient *resty.Client
	gqlClient   *resty.Client
	authUrl string
}

func NewTwitchService(helixClientId string, helixSecret string, gqlClientId string) *TwitchService {
	return &TwitchService{
		helixClient: resty.New().
			SetTimeout(10 * time.Second).
			SetBaseURL(baseUrl).
			SetDebug(true).
			OnBeforeRequest(addAuthHeaderMiddleware(helixClientId, helixSecret)),
		gqlClient:  resty.New().
			SetTimeout(10 * time.Second).
			SetBaseURL(gqlUrl).
			SetDebug(true).
			SetHeader("Client-ID", gqlClientId),
		authUrl: oauthBaseURL,
	}
}
