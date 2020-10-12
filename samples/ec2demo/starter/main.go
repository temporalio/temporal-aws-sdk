package main

import (
	"context"
	"go.temporal.io/aws-sdk/samples/ec2demo"
	"go.temporal.io/sdk/client"
	"log"
)

const taskQueue = "ec2demo"

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		TaskQueue: taskQueue,
	}
	wRun, err := c.ExecuteWorkflow(context.Background(), options, ec2demo.KeepInstance)
	if err != nil {
		log.Fatalln("Failure starting workflow", err)
	}
	log.Printf("Started KeepInstanceWorkflow. WorkflowID=%v", wRun.GetID())
}
