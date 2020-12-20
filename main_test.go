package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

const req1 string = `
{
    "requestContext": {
        "elb": {
            "targetGroupArn": "arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/lambda-279XGJDqGZ5rsrHC2Fjr/49e9d65c45c6791a"
        }
    },
    "httpMethod": "GET",
    "path": "/lambda",
    "queryStringParameters": {
        "query": "1234ABCD"
    },
    "headers": {
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
        "accept-encoding": "gzip",
        "accept-language": "en-US,en;q=0.9",
        "connection": "keep-alive",
        "host": "lambda-alb-123578498.us-east-2.elb.amazonaws.com",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36",
        "x-amzn-trace-id": "Root=1-5c536348-3d683b8b04734faae651f476",
        "x-forwarded-for": "72.12.164.125",
        "x-forwarded-port": "80",
        "x-forwarded-proto": "http",
        "x-imforwards": "20"
    },
    "body": "",
    "isBase64Encoded": false
}
`

func TestGetClientIP(t *testing.T) {
	var request events.ALBTargetGroupRequest
	json.Unmarshal([]byte(req1), &request)

	clientip := getClientIP(request)
	if clientip != "72.12.164.125" {
		t.Errorf("getClientIP returned %s", clientip)
	}
}

func TestHandleRequest(t *testing.T) {
	var request events.ALBTargetGroupRequest
	json.Unmarshal([]byte(req1), &request)

	response, err := HandleRequest(context.Background(), request)
	if err != nil {
		t.Errorf(" returned error: %s", err)
	}
	if response.Body != "72.12.164.125" {
		t.Errorf(" returned body: %s", response.Body)
	}
}
