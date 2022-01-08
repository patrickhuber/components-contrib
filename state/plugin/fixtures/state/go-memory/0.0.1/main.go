package main

import (
	"encoding/json"
	"fmt"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/components-contrib/state/plugin"
	goplugin "github.com/hashicorp/go-plugin"
)

type Store struct {
	data map[string][]byte
}

func (s *Store) Init(metadata state.Metadata) error {
	s.data = map[string][]byte{}
	return nil
}

func (s *Store) Features() []state.Feature {
	return []state.Feature{}
}

func (s *Store) Delete(req *state.DeleteRequest) error {
	delete(s.data, req.Key)
	return nil
}

func (s *Store) Get(req *state.GetRequest) (*state.GetResponse, error) {
	value, ok := s.data[req.Key]
	if !ok {
		return nil, fmt.Errorf("missing key '%s' in memory store", req.Key)
	}

	return &state.GetResponse{
		Metadata: req.Metadata,
		Data:     value,
	}, nil
}

func (s *Store) Set(req *state.SetRequest) error {
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

	s.data[req.Key] = bytes

	return nil
}

func (s *Store) Ping() error {
	return nil
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
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: plugin.Handshake,
		Plugins: map[string]goplugin.Plugin{
			// TODO: check if this should be plugin.RPC
			plugin.ProtocolRPC: &plugin.GRPCStatePlugin{
				Impl: &Store{},
			},
		},
		GRPCServer: goplugin.DefaultGRPCServer,
	})
}
