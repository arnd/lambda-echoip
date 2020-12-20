# Lambda funtion answer to: What's my IP?

When invoked via an AWS application level loadbalancer (ELB) the function will return the IP of the requesting client.

## Prerequisites

* aws cli installed and configured with credentials and default region
* GNU Make
* Go >= 1.15.x

## Install

1. Create function execution role and note down the Arn of the role (e.g. `arn:aws:iam::123456789012:role/lambda-echoip`)
```shell
aws iam create-role --role-name lambda-echoip --assume-role-policy-document file://trust-policy.json`
```
2. Build the deployment package `function.zip`:
```
make build
```
3. Create function and note down the Function ARN (e.g. `arn:aws:lambda:eu-west-1:123456789012:function:lambda-echoip`):
```shell
aws lambda create-function --function-name lambda-echoip --runtime go1 --handler lambda-echoip.handler --zip-file fileb://function.zip --role arn:aws:iam::123456789012:role/lambda-echoip`
```
4. Create ELB target group for the function, and note down the TargetGroupArn (e.g. `arn:aws:elasticloadbalancing:eu-west-1:123456789012:targetgroup/lambda-echoip/7e8a3b1bb81b9338"`)
```shell
aws elbv2 create-target-group --name lambda-echoip --target-type lambda
```
5. Grant ELB invocation rights on function:
```shell
aws lambda add-permission \
    --function-name lambda-echoip \
    --statement-id load-balancer \
    --principal elasticloadbalancing.amazonaws.com \
    --action lambda:InvokeFunction \
    --source-arn arn:aws:elasticloadbalancing:eu-west-1:123456789012:targetgroup/lambda-echoip/7e8a3b1bb81b9338
```
6. Register function as ELB target:
```shell
aws elbv2 register-targets \
    --target-group-arn arn:aws:elasticloadbalancing:eu-west-1:123456789012:targetgroup/lambda-echoip/7e8a3b1bb81b9338 \
    --targets Id=arn:aws:lambda:eu-west-1:123456789012:function:lambda-echoip
```
7. Test function:
```shell
curl echoip-123456789.eu-west-1.elb.amazonaws.com
```

## Development

* `make help` shows Makefile targets
* `make test`  runs the tests
* `make deploy` will update an existing deployment with the new code
