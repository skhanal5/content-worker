package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

var ActivityOptions = workflow.ActivityOptions{
	StartToCloseTimeout: 10 * time.Second,
}
