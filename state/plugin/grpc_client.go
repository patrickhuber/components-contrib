package plugin

import (
	"context"
	"encoding/json"

	"github.com/dapr/components-contrib/state"
)

// see https://developers.google.com/protocol-buffers/docs/reference/go-generated
// see https://github.com/src-d/proteus

// GRPCClient provides a grpc client for the state store
type GRPCClient struct {
	client StateClient
}

func NewGRPCClient() state.Store {
	return &GRPCClient{}
}

func (c *GRPCClient) Features() []state.Feature {
	return nil
}
func (c *GRPCClient) Init(req state.Metadata) error {
	metadata := &Metadata{
		Properties: map[string]string{},
	}
	for k, v := range req.Properties {
		metadata.Properties[k] = v
	}
	_, err := c.client.Init(context.Background(), metadata)
	return err
}

func (c *GRPCClient) Get(req *state.GetRequest) (*state.GetResponse, error) {
	request := &GetRequest{}
	response, err := c.client.Get(context.Background(), request)
	if err != nil {
		return nil, err
	}

	etag := response.GetEtag()
	return &state.GetResponse{
		Data:     response.GetData(),
		ETag:     &etag,
		Metadata: response.GetMetadata(),
	}, nil
}

func (c *GRPCClient) Set(req *state.SetRequest) error {
	var bytes []byte
	byteArray, isBinary := req.Value.([]uint8)
	if isBinary {
		bytes = byteArray
	} else {
		marshalled, err := json.Marshal(req.Value)
		if err != nil {
			return err
		}
		bytes = marshalled
	}

	request := &SetRequest{
		Key:      req.GetKey(),
		Value:    bytes,
		Etag:     req.GetKey(),
		Metadata: req.GetMetadata(),
		Options: &SetStateOption{
			Concurrency: req.Options.Concurrency,
			Consistency: req.Options.Consistency,
		},
	}
	_, err := c.client.Set(context.Background(), request)
	return err
}

func (c *GRPCClient) Ping() error {
	empty := &Empty{}
	_, err := c.client.Ping(context.Background(), empty)
	return err
}

func (c *GRPCClient) Delete(req *state.DeleteRequest) error {
	request := &DeleteRequest{
		Key:      req.GetKey(),
		Etag:     *req.ETag,
		Metadata: req.GetMetadata(),
		Options: &DeleteStateOption{
			Concurrency: req.Options.Concurrency,
			Consistency: req.Options.Consistency,
		},
	}
	_, err := c.client.Delete(context.Background(), request)
	return err
}

func (c *GRPCClient) BulkDelete(req []state.DeleteRequest) error {
	return nil
}

func (c *GRPCClient) BulkGet(req []state.GetRequest) (bool, []state.BulkGetResponse, error) {
	return false, nil, nil
}

func (c *GRPCClient) BulkSet(req []state.SetRequest) error {
	return nil
}
