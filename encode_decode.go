package nats

import "github.com/nats-io/go-nats"

// DecodeRequestFunc extracts a user-domain request object from a gRPC request.
// It's designed to be used in gRPC servers, for server-side endpoints. One
// straightforward DecodeRequestFunc could be something that decodes from the
// gRPC request message to the concrete request type.
type DecodeRequestFunc func(msg *nats.Msg) (request interface{}, err error)

// EncodeRequestFunc encodes the passed request object into the gRPC request
// object. It's designed to be used in gRPC clients, for client-side endpoints.
// One straightforward EncodeRequestFunc could something that encodes the object
// directly to the gRPC request message.
type EncodeRequestFunc func(msg interface{}) (payload []byte, err error)

// EncodeResponseFunc encodes the passed response object to the gRPC response
// message. It's designed to be used in gRPC servers, for server-side endpoints.
// One straightforward EncodeResponseFunc could be something that encodes the
// object directly to the gRPC response message.
type EncodeResponseFunc func(response interface{}) (payload []byte, err error)

// DecodeResponseFunc extracts a user-domain response object from a gRPC
// response object. It's designed to be used in gRPC clients, for client-side
// endpoints. One straightforward DecodeResponseFunc could be something that
// decodes from the gRPC response message to the concrete response type.
type DecodeResponseFunc func(msg *nats.Msg) (response interface{}, err error)
