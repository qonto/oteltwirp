package oteltwirp

import (
	"go.opentelemetry.io/otel/attribute"
)

const instrumentationName = "github.com/qonto/oteltwirp"

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
