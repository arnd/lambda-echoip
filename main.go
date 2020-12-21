package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getClientIP(request events.ALBTargetGroupRequest) string {
	fwdHeader := request.Headers["x-forwarded-for"]
	// if there are multiple IPs, use the first in chain
	IPs := strings.Split(fwdHeader, ", ")
	clientIP := IPs[0]
	return clientIP
}

// HandleRequest will handle request event of Application Loadbalancer (ELB)
// It returns the httpResponse
func HandleRequest(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	httpResponse := events.ALBTargetGroupResponse{
		Body:              getClientIP(request),
		StatusCode:        200,
		StatusDescription: "200 OK",
		IsBase64Encoded:   false,
		Headers:           map[string]string{}}

	return httpResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
