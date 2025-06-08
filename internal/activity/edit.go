package activity

import (
	"clip-farmer-workflow/internal/service/edit"
	"context"
	"fmt"
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

	strategy,err := a.getStrategy(input.Strategy)
	
	if err != nil {
		return err
	}

	err = a.EditManager.Render(input.InputPath, input.OutputPath, strategy)
	
	if err != nil {
		return err
	}
	return nil
}