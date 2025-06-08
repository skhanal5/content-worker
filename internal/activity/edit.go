package activity

import (
	"clip-farmer-workflow/internal/service/edit"
	"context"
	"fmt"
	"time"
	"go.temporal.io/sdk/activity"
)

type VideoStrategyType string

const (
	BlurredOverlay VideoStrategyType = "blurred_overlay"
	Resize         VideoStrategyType = "resize"
	Watermark      VideoStrategyType = "watermark"
)

type EditVideoInput struct {
	InputPath string
	OutputPath string
	Strategy VideoStrategyType
}


func (a Activity) getStrategy(strategyType VideoStrategyType) (edit.EditingStrategy, error) {
	switch strategyType {
	case BlurredOverlay:
		return &edit.BlurredOverlayStrategy{}, nil
	default:
		return nil, fmt.Errorf("unsupported strategy")
	}
}

func (a *Activity) EditVideo(ctx context.Context, input EditVideoInput) (error) {
    logger := activity.GetLogger(ctx)
	strategy,err := a.getStrategy(input.Strategy)
	logger.Info("Using strategy:", strategy)

	stop := make(chan struct{})
    go func() {
        ticker := time.NewTicker(10 * time.Second)
        defer ticker.Stop()
        for {
            select {
            case <-ticker.C:
                activity.RecordHeartbeat(ctx)
                logger.Info("Heartbeat sent during render")
            case <-stop:
                return
            }
        }
    }()
	
	if err != nil {
		return err
	}

	err = a.EditManager.Render(input.InputPath, input.OutputPath, strategy)
	close(stop)

	if err != nil {
		logger.Error("Render failed", "error", err)
		return err
	}

	logger.Info("Video editing complete")
	return nil
}