package activity

import "context"

type DownloadClipInput struct {
	ID  string
	URL string
}


func (a *Activity) DownloadClip(ctx context.Context, input DownloadClipInput) (string, error) {
	output, err := a.DownloadManager.Download(input.URL, input.ID)
	if err != nil {
		return "", err
	}
	return output, nil
}