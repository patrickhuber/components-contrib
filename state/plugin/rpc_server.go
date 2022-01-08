package plugin

import "github.com/dapr/components-contrib/state"

// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// real implementation
	Impl state.Store
}

func (s *RPCServer) Init(metadata *state.Metadata, resp *interface{}) error {
	return s.Impl.Init(*metadata)
}

func (s *RPCServer) Features(req *interface{}, features *[]state.Feature) error {
	*features = s.Impl.Features()
	return nil
}

func (s *RPCServer) Delete(req *state.DeleteRequest, resp *interface{}) error {
	return s.Impl.Delete(req)
}

func (s *RPCServer) Get(req *state.GetRequest, resp *state.GetResponse) error {
	response, err := s.Impl.Get(req)
	if err != nil {
		return err
	}
	*resp = *response
	return nil
}

func (s *RPCServer) Set(req *state.SetRequest, resp *interface{}) error {
	return s.Impl.Set(req)
}

func (s *RPCServer) Ping(args *interface{}, resp *interface{}) error {
	return s.Impl.Ping()
}
