package oteltwirp

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

const instrumentationName = "github.com/qonto/twirp-otel"

// Semantic conventions for attribute keys for twirp.
const (
	// Type of message transmitted or received.
	RPCMessageTypeKey = attribute.Key("message.type")
)

// Semantic conventions for common RPC attributes.
var (
	// Semantic conventions for RPC message types.
	RPCMessageTypeSent     = RPCMessageTypeKey.String("SENT")     //nolint:gochecknoglobals
	RPCMessageTypeReceived = RPCMessageTypeKey.String("RECEIVED") //nolint:gochecknoglobals
)
