package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// this is just for debugging
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("My request: %+v", request)
	fmt.Printf("My context: %+v", ctx)

	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadGateway,
			Body:       "Something went wrong :(",
		}, nil
	}

	out := &struct {
		Headers       map[string]string `json:"headers,omitempty"`
		ClientContext struct {
			ClientInfo struct {
				AppPackageName string `json:"app_package_name,omitempty"`
				AppTitle       string `json:"app_title,omitempty"`
				AppVersionCode string `json:"app_version_code,omitempty"`
				InstallationID string `json:"installation_id,omitempty"`
			} `json:"client_info,omitemtpy,omitempty"`
			Custom map[string]string `json:"custom,omitempty"`
			Env    map[string]string `json:"env,omitempty"`
		} `json:"client_context,omitempty"`
		AWSInfo struct {
			AwsRequestID       string `json:"aws_request_id,omitempty"`
			InvokedFunctionArn string `json:"invoked_function_arn,omitempty"`
		} `json:"aws_info,omitempty"`
		Netlify map[string]interface{} `json:"netlify,omitempty"`
	}{}

	out.Headers = request.Headers
	out.ClientContext.Env = lc.ClientContext.Env
	out.ClientContext.Custom = lc.ClientContext.Custom
	out.ClientContext.ClientInfo.AppPackageName = lc.ClientContext.Client.AppPackageName
	out.ClientContext.ClientInfo.AppTitle = lc.ClientContext.Client.AppTitle
	out.ClientContext.ClientInfo.AppVersionCode = lc.ClientContext.Client.AppVersionCode
	out.ClientContext.ClientInfo.InstallationID = lc.ClientContext.Client.InstallationID
	out.AWSInfo.AwsRequestID = lc.AwsRequestID
	out.AWSInfo.InvokedFunctionArn = lc.InvokedFunctionArn

	// the netlify struct is encoded
	nfstr, ok := lc.ClientContext.Custom["netlify"]
	if ok {
		fmt.Println("Found netlify context")
		decoded, err := base64.StdEncoding.DecodeString(nfstr)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       err.Error(),
			}, nil
		}
		parsed := make(map[string]interface{})
		if err := json.Unmarshal(decoded, &parsed); err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       err.Error(),
			}, nil
		}
		out.Netlify = parsed
	}

	fmt.Printf("Marshaling the output: %+v\n", out)
	bs, err := json.Marshal(out)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	fmt.Println("It all seems good")
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bs),
	}, nil
}

func main() {
	lambda.Start(handler)
}
