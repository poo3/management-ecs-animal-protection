package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

const desiredCount int64  = 0

var cluster = "animal-protection-cluster-frontend"

func main() {
	lambda.Start(stopServies)
}

func stopServies() {
	services := []string{
		"animal-protection-service-backend",
		"animal-protection-service-frontend",
	}

	for _, service := range services {
		svc := ecs.New(session.New())
		input := &ecs.UpdateServiceInput{
			Cluster:        aws.String(cluster),
			Service:        aws.String(service),
			DesiredCount:   aws.Int64(desiredCount),
		}
		result, err := svc.UpdateService(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				fmt.Println(aerr.Code(), aerr.Error())
			}
			return
		}
		fmt.Println(result)
	}
}