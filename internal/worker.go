package internal

import (
	"clip-farmer-workflow/internal/activity"
	"clip-farmer-workflow/internal/workflow"
	"log"

    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
)

func registerWorker(client client.Client) (worker.Worker, error) {	
	worker := worker.New(client, "TASK QUEUE NAME", worker.Options{})
	worker.RegisterWorkflow(workflow.HelloWorldWorkflow)
	worker.RegisterActivity(activity.HelloWorldActivity)
	return worker, nil
}

func StartWorker(client client.Client) {
	worker, err := registerWorker(client)
	if err != nil {
		log.Fatalln("Unable to register Worker", err)
	}
	err = worker.Start()
    if err != nil {
        log.Fatalln("Unable to start Worker", err)
    }
}