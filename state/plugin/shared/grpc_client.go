package shared

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/components-contrib/state/plugin/proto"
	"github.com/dapr/components-contrib/state/utils"
)

// see https://developers.google.com/protocol-buffers/docs/reference/go-generated
// see https://github.com/src-d/proteus

// GRPCClient provides a grpc client for the state store
type GRPCClient struct {
	client proto.StateClient
}

func NewGRPCClient(client proto.StateClient) state.Store {
	return &GRPCClient{
		client: client,
	}
}

func (c *GRPCClient) Features() []state.Feature {
	return nil
}

func (c *GRPCClient) Init(req state.Metadata) error {
	metadata := &proto.Metadata{
		Properties: map[string]string{},
	}
	for k, v := range req.Properties {
		metadata.Properties[k] = v
	}
	_, err := c.client.Init(context.Background(), metadata)
	return err
}

func (c *GRPCClient) Get(req *state.GetRequest) (*state.GetResponse, error) {

	request := &proto.GetRequest{
		Key:      req.Key,
		Metadata: req.Metadata,
		Options: &proto.GetStateOption{
			Consistency: "",
		},
	}

	etag := ""
	emptyResponse := &state.GetResponse{
		ETag:     &etag,
		Metadata: map[string]string{},
		Data:     []byte{},
	}

	response, err := c.client.Get(context.Background(), request)
	if err != nil {
		return emptyResponse, err
	}
	if response == nil {
		return emptyResponse, fmt.Errorf("response is nil")
	}

	etag = response.GetEtag()
	return &state.GetResponse{
		Data:     response.GetData(),
		ETag:     &etag,
		Metadata: response.GetMetadata(),
	}, nil
}

func (c *GRPCClient) Set(req *state.SetRequest) error {
	var bytes []byte
	switch t := req.Value.(type) {
	case string:
		bytes = []byte(t)
	case []byte:
		bytes = t
	default:
		var err error
		if bytes, err = utils.Marshal(t, json.Marshal); err != nil {
			return err
		}
	}
	request := &proto.SetRequest{
		Key:      req.GetKey(),
		Value:    bytes,
		Etag:     req.GetKey(),
		Metadata: req.GetMetadata(),
		Options: &proto.SetStateOption{
			Concurrency: req.Options.Concurrency,
			Consistency: req.Options.Consistency,
		},
	}
	_, err := c.client.Set(context.Background(), request)
	return err
}

func (c *GRPCClient) Ping() error {
	empty := &proto.Empty{}
	_, err := c.client.Ping(context.Background(), empty)
	return err
}

func (c *GRPCClient) Delete(req *state.DeleteRequest) error {
	request := &proto.DeleteRequest{
		Key:      req.GetKey(),
		Etag:     *req.ETag,
		Metadata: req.GetMetadata(),
		Options: &proto.DeleteStateOption{
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
