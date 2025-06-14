package download

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)


type DownloadManager interface {
	Download(url string, filepath string) error
}

type DownloadService struct {
	client *resty.Client
}

func (d *DownloadService) Download(url string, filepath string) error {
	if url == "" {
        return fmt.Errorf("url cannot be empty")
    }
	
	resp, err := d.client.R().
	SetOutput(filepath).
	Get(url)

	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("failed to download file: %s", resp.Status())
	}
	return nil
}

func NewDownloadService() *DownloadService {
	return &DownloadService{
		client: resty.New().
		SetTimeout(30 * time.Second),
	}
}