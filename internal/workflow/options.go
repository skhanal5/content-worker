package workflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var retryPolicy = &temporal.RetryPolicy{
	InitialInterval:    1 * time.Second,
	BackoffCoefficient: 2.0,
	MaximumInterval:   100 * time.Second,
	MaximumAttempts:     1,
}

var ActivityOptions = workflow.ActivityOptions{
	StartToCloseTimeout: 10 * time.Second,
	RetryPolicy: 	   retryPolicy,
}
