package shared

import (
	"context"
	"fmt"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/components-contrib/state/plugin/proto"
)

type GRPCServer struct {
	proto.UnimplementedStateServer
	// this is the real implementation
	Impl state.Store
}

func (s *GRPCServer) Init(ctx context.Context, req *proto.Metadata) (*proto.Empty, error) {
	metadata := state.Metadata{
		Properties: req.GetProperties(),
	}
	return &proto.Empty{}, s.Impl.Init(metadata)
}

func (s *GRPCServer) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	request := &state.GetRequest{
		Key:      req.GetKey(),
		Metadata: req.GetMetadata(),
		Options: state.GetStateOption{
			Consistency: req.GetOptions().GetConsistency(),
		},
	}
	response, err := s.Impl.Get(request)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, fmt.Errorf("response is nil")
	}
	etag := ""
	if response.ETag != nil {
		etag = *response.ETag
	}
	return &proto.GetResponse{
		Data:     response.Data,
		Etag:     etag,
		Metadata: response.Metadata,
	}, nil
}

func (s *GRPCServer) Set(ctx context.Context, req *proto.SetRequest) (*proto.Empty, error) {
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
	return &proto.Empty{}, err
}

func (s *GRPCServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.Empty, error) {
	deleteRequest := &state.DeleteRequest{
		Key:      req.Key,
		ETag:     &req.Etag,
		Metadata: req.Metadata,
		Options: state.DeleteStateOption{
			Concurrency: req.Options.Concurrency,
			Consistency: req.Options.Consistency,
		},
	}
	return &proto.Empty{}, s.Impl.Delete(deleteRequest)
}

func (s *GRPCServer) Ping(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, s.Impl.Ping()
}
