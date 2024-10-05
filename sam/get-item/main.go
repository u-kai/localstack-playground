package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type request struct {
	Id string `json:"id"`
}

type response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
}

func handler(ctx context.Context, req request) (response, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return response{
			StatusCode: 500,
		}, err
	}
	client := dynamodb.NewFromConfig(cfg)

	tableName := os.Getenv("TABLE_NAME")
	item, err := client.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: aws.String(tableName),
			Key: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{Value: req.Id},
			},
		},
	)
	if err != nil {
		return response{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}
	data := item.Item["name"].(*types.AttributeValueMemberS).Value
	return response{
		Body:       data,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
