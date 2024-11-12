package oteltwirp

import (
	"context"
	"fmt"

	"github.com/twitchtv/twirp"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// spanInfo returns a span name and all appropriate attributes from the context
func spanInfo(ctx context.Context) (string, []attribute.KeyValue) {
	packageName, _ := twirp.PackageName(ctx)
	serviceName, _ := twirp.ServiceName(ctx)
	methodName, _ := twirp.MethodName(ctx)
	attr := []attribute.KeyValue{
		semconv.RPCServiceKey.String(serviceName),
		semconv.RPCMethodKey.String(methodName),
	}
	spanName := fmt.Sprintf("%s.%s/%s", packageName, serviceName, methodName)

	return spanName, attr
}
