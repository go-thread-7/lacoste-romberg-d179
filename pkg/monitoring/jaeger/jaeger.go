package otel

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

type JaegerConfig struct {
	Server      string `mapstructure:"server"`
	ServiceName string `mapstructure:"serviceName"`
	TracerName  string `mapstructure:"tracerName"`
}

func TracerProvider(ctx context.Context, cfg *JaegerConfig) (trace.Tracer, error) {
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	env := os.Getenv("APP_ENV")
	if env != "production" {
		env = "development"
	}

	// Create the tracer provider with the OTLP exporter
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.ServiceName),
			attribute.String("environment", env),
		)),
	)

	go func() {
		for {
			select {
			case <-ctx.Done():
				err = tp.Shutdown(ctx)
				fmt.Println("open-telemetry exited properly")
				if err != nil {
					return
				}
			}
		}
	}()

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))

	t := tp.Tracer(cfg.TracerName)

	return t, nil
}
