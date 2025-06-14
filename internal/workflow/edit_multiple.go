package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"fmt"
	"log"
	"os"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Video struct {
	InputPath string `json:"input_path,required"`
	Title string `json:"title,required"`
	Strategy string `json:"strategy,required"`
}

type EditMultipleWorkflowInput struct {
	OutputDirectory string `json:"output_directory,required"`
	Videos []Video `json:"videos,required"`
}


func EditMultipleWorkflow(ctx workflow.Context, input EditMultipleWorkflowInput) error {
	
	ao := workflow.ActivityOptions{
		RetryPolicy: retryPolicy,
        StartToCloseTimeout:    time.Second * 180,
        HeartbeatTimeout:       time.Second * 15,
	}
	
	ctx = workflow.WithActivityOptions(ctx, ao)
	var a activity.Activity
	log.Printf("Kicking off Edit Multiple Workflow with payload: %s", input)
	
	outputDir := input.OutputDirectory
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Printf("Cannot find the output directory: %v", err)
		return err
	}
	for _, video := range input.Videos {
		outputPath := fmt.Sprintf("%s/%s.mp4", outputDir, video.Title)
		editInput := activity.EditVideoInput{
			InputPath: video.InputPath,
			OutputPath:  outputPath,
			Strategy: activity.VideoStrategyType(video.Strategy),
			Title: video.Title,
		}

		err := workflow.ExecuteActivity(ctx, a.EditVideo, editInput).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	log.Println("Completed Edit Multiple Workflow")
	return nil
}