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
	Name string `json:"name"`
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
	id := "id_" + req.Name
	_, err = client.PutItem(
		ctx,
		&dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item: map[string]types.AttributeValue{
				"id":   &types.AttributeValueMemberS{Value: id},
				"name": &types.AttributeValueMemberS{Value: req.Name},
			},
		},
	)
	if err != nil {
		return response{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	return response{
		Body:       "id:" + id,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
