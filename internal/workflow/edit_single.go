package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"fmt"
	"log"
	"os"
	"time"

	"go.temporal.io/sdk/workflow"
)

type EditSingleWorkflowInput struct {
	InputPath       string `json:"input_path"`
	OutputDirectory string `json:"output_directory"`
	Strategy        string `json:"strategy"`
	Title           string `json:"title"`
}

func EditSingleWorkflow(ctx workflow.Context, input EditSingleWorkflowInput) error {

	ao := workflow.ActivityOptions{
		RetryPolicy:         retryPolicy,
		StartToCloseTimeout: time.Second * 180,
		HeartbeatTimeout:    time.Second * 15,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	var a activity.Activity
	log.Printf("Kicking off Single Edit Workflow with payload: %s", input)

	outputDir := input.OutputDirectory
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Printf("Cannot find the output directory: %v", err)
		return err
	}

	outputPath := fmt.Sprintf("%s/%s.mp4", outputDir, input.Title)

	editInput := activity.EditVideoInput{
		InputPath:  input.InputPath,
		OutputPath: outputPath,
		Style:      activity.EditStyle(input.Strategy),
		Title:      input.Title,
	}

	err = workflow.ExecuteActivity(ctx, a.EditVideo, editInput).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, a.DeleteTmpVideo, input.InputPath).Get(ctx, nil)
	if err != nil {
		return err
	}

	log.Println("Completed Single Edit Workflow")
	return nil
}
