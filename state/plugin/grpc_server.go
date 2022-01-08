package plugin

import (
	"context"

	"github.com/dapr/components-contrib/state"
)

type GRPCServer struct {
	// this is the real implementation
	Impl state.Store
}

func (s *GRPCServer) Init(ctx context.Context, req *Metadata) (*Empty, error) {
	metadata := state.Metadata{
		Properties: req.GetProperties(),
	}
	return &Empty{}, s.Impl.Init(metadata)
}

func (s *GRPCServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	getRequest := &state.GetRequest{
		Key:      req.Key,
		Metadata: req.Metadata,
		Options: state.GetStateOption{
			Consistency: req.Options.Consistency,
		},
	}
	getResponse, err := s.Impl.Get(getRequest)
	if err != nil {
		return nil, err
	}
	response := &GetResponse{
		Data:     getResponse.Data,
		Etag:     *getResponse.ETag,
		Metadata: getResponse.Metadata,
	}
	return response, nil
}

func (s *GRPCServer) Set(ctx context.Context, req *SetRequest) (*Empty, error) {
	setRequest := &state.SetRequest{
		Key:   req.Key,
		ETag:  &req.Etag,
		Value: req.Value,
		Options: state.SetStateOption{
			Concurrency: req.Options.Concurrency,
			Consistency: req.Options.Consistency,
		},
	}
	err := s.Impl.Set(setRequest)
	return &Empty{}, err
}

func (s *GRPCServer) Delete(ctx context.Context, req *DeleteRequest) (*Empty, error) {
	deleteRequest := &state.DeleteRequest{
		Key:      req.Key,
		ETag:     &req.Etag,
		Metadata: req.Metadata,
		Options: state.DeleteStateOption{
			Concurrency: req.Options.Concurrency,
			Consistency: req.Options.Consistency,
		},
	}
	return &Empty{}, s.Impl.Delete(deleteRequest)
}

func (s *GRPCServer) Ping(ctx context.Context, req *Empty) (*Empty, error) {
	return &Empty{}, s.Impl.Ping()
}
