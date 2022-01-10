package main

import (
	"encoding/json"
	"fmt"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/components-contrib/state/plugin/proto"
	"github.com/dapr/components-contrib/state/plugin/shared"
	"github.com/dapr/components-contrib/state/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type Store struct {
	data   map[string][]byte
	logger hclog.Logger
	proto.UnimplementedStateServer
}

func (s *Store) Init(metadata state.Metadata) error {
	s.Log("Init(metadata)")
	for k := range s.data {
		delete(s.data, k)
	}
	return nil
}

func (s *Store) Features() []state.Feature {
	s.Log("Features()")
	return []state.Feature{}
}

func (s *Store) Delete(req *state.DeleteRequest) error {
	s.Log("Delete(%s)", req.Key)
	delete(s.data, req.Key)
	return nil
}

func (s *Store) Get(req *state.GetRequest) (*state.GetResponse, error) {
	emptyResponse := &state.GetResponse{
		Data:     nil,
		ETag:     nil,
		Metadata: nil,
	}
	if req == nil {
		s.Log("Get(<nil>)")
		return emptyResponse, nil
	}

	value, ok := s.data[req.Key]
	if !ok {
		s.Log("Get(%s) not found", req.Key)
		return emptyResponse, nil
	}

	metadata := map[string]string{}
	for k, v := range req.Metadata {
		metadata[k] = v
	}

	etag := ""

	return &state.GetResponse{
		Data:     value,
		ETag:     &etag,
		Metadata: map[string]string{},
	}, nil
}

func (s *Store) Set(req *state.SetRequest) error {
	s.Log("Set(%s)", req.Key)
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

	s.data[req.Key] = bytes

	return nil
}

func (s *Store) Ping() error {
	s.Log("Ping()")
	return nil
}

func (s *Store) Log(message string, args ...interface{}) {
	if s.logger == nil {
		return
	}
	s.logger.Debug("go-memory: "+message, args)
}

func (s *Store) Error(message string, args ...interface{}) error {
	err := fmt.Errorf(message, args...)
	if s.logger == nil {
		return err
	}
	s.logger.Error("go-memory: "+message, args)
	return err
}

func (s *Store) BulkDelete(req []state.DeleteRequest) error {
	return nil
}

func (s *Store) BulkGet(req []state.GetRequest) (bool, []state.BulkGetResponse, error) {
	return false, nil, nil
}

func (s *Store) BulkSet(req []state.SetRequest) error {
	return nil
}

func main() {
	logger := hclog.Default()
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			// should this be GRPC?
			shared.ProtocolGRPC: &shared.GRPCStatePlugin{
				Impl: &Store{
					data:   map[string][]byte{},
					logger: logger,
				},
			},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
