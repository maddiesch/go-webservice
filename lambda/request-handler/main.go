package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdaurl"
	"github.com/maddiesch/go-webservice/internal/injector"
	sloglambda "github.com/maddiesch/slog-lambda"
	"github.com/samber/do/v2"
)

func init() {
	lambdaLogHandler := sloglambda.NewHandler(os.Stdout)
	loggerInstance := slog.New(lambdaLogHandler)
	slog.SetDefault(loggerInstance)
}

func main() {
	i := injector.New()

	requestHandler := createRequestHandler(i)

	// Wrap the request handler with the lambdaurl package to map the request/response types
	handler := lambdaurl.Wrap(requestHandler)

	// Ensure the injector resources are cleaned up properly when the Lambda function is shutdown
	shutdown := injector.LambdaShutdownOption(i)

	lambda.StartWithOptions(handler, shutdown)
}

func createRequestHandler(do.Injector) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
}
