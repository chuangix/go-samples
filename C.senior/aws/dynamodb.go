package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	resp, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}

	resp1, err1 := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("st-nds-db-smart-nja2-22000400-all"),
		Limit:     aws.Int32(10),
	})
	if err1 != nil {
		log.Fatalf("failed to list tables, %v", err1)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp1.Items {
		fmt.Println(tableName)
	}
	fmt.Println(resp1.LastEvaluatedKey)

	resp2, err2 := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:         aws.String("st-nds-db-smart-nja2-22000400-all"),
		Limit:             aws.Int32(5),
		ExclusiveStartKey: resp1.LastEvaluatedKey,
	})
	if err2 != nil {
		log.Fatalf("failed to list tables, %v", err2)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp2.Items {
		fmt.Println(tableName)
	}
	fmt.Println(resp2.LastEvaluatedKey)
}
