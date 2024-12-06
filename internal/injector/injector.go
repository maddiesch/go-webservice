package injector

import (
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/samber/do/v2"
)

type Option func(do.Injector)

func WithOverride[T any](provider do.Provider[T]) Option {
	return func(i do.Injector) {
		do.Override(i, provider)
	}
}

func WithOverrideValue[T any](value T) Option {
	return func(i do.Injector) {
		do.OverrideValue(i, value)
	}
}

func New(options ...Option) do.Injector {
	i := do.New()

	for _, option := range options {
		option(i)
	}

	return i
}

func LambdaShutdownOption(i do.Injector, errors ...func(error)) lambda.Option {
	if len(errors) == 0 {
		errors = append(errors, func(err error) {
			slog.Error("Failed to shutdown injector with error", slog.String("error", err.Error()))
		})
	}

	return lambda.WithEnableSIGTERM(func() {
		slog.Debug("Shutting down injector, Lambda received SIGTERM")

		err := i.Shutdown()

		if err != nil {
			for _, e := range errors {
				e(err)
			}
		}
	})
}
