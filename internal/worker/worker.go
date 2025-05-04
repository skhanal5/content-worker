package worker

import (
	"clip-farmer-workflow/internal/activity"
	"clip-farmer-workflow/internal/config"
	"clip-farmer-workflow/internal/workflow"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func registerWorker(client client.Client, cfg config.Config) (worker.Worker, error) {
	services := activity.NewActivity(cfg)
	worker := worker.New(client, "default", worker.Options{})
	worker.RegisterWorkflow(workflow.HelloWorldWorkflow)
	worker.RegisterActivity(services)
	return worker, nil
}

func StartWorker(cfg config.Config) {
	c, err := client.Dial(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: client.DefaultNamespace,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w, err := registerWorker(c, cfg)
	if err != nil {
		log.Fatalln("Unable to register Worker", err)
	}
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}
}
