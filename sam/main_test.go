package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func TestInvokeLambda(t *testing.T) {
	ctx := context.TODO()
	name := os.Args[1]
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	client := lambda.NewFromConfig(cfg, func(o *lambda.Options) {
		o.BaseEndpoint = aws.String("http://localhost:4566")
	})
	res, err := client.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String("put-item"),
		Payload:      []byte(fmt.Sprintf(`{"name":"%s"}`, name)),
	})
	if err != nil {
		panic("unable to invoke lambda, " + err.Error())
	}
	fmt.Println(string(res.Payload))

	res, err = client.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String("get-item"),
		Payload:      []byte(fmt.Sprintf(`{"id":"id_%s"}`, name)),
	})
	if err != nil {
		panic("unable to invoke lambda, " + err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("unexpected status code: %d", res.StatusCode)
	}
	type payload struct {
		Body       string `json:"body"`
		StatusCode int    `json:"statusCode"`
	}
	var p payload
	if err := json.Unmarshal([]byte(res.Payload), &p); err != nil {
		t.Errorf("unable to unmarshal payload, %v", err)
	}
	if p.Body != name {
		t.Errorf("unexpected body: %s", p.Body)
	}
	if p.StatusCode != 200 {
		t.Errorf("unexpected status code: %d", p.StatusCode)
	}

}
