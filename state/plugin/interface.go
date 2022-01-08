package plugin

import (
	"context"
	"net/rpc"

	"github.com/dapr/components-contrib/state"
	"github.com/hashicorp/go-plugin"
	goplugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// see https://github.com/hashicorp/go-plugin/blob/master/examples/grpc/shared/interface.go

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = goplugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "76d3865e-360a-416a-bdf3-7f9891a4a2b8",
}

const (
	ProtocolRPC  = "state"
	ProtocolGRPC = "state_grpc"
)

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]goplugin.Plugin{
	ProtocolRPC:  &RPCStatePlugin{},
	ProtocolGRPC: &GRPCStatePlugin{},
}

type GRPCStatePlugin struct {
	goplugin.Plugin
	Impl state.Store
}

func (p *GRPCStatePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterStateServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GRPCStatePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewStateClient(c)}, nil
}

type RPCStatePlugin struct {
	Impl state.Store
}

func (p *RPCStatePlugin) Server(b *goplugin.MuxBroker) (interface{}, error) {
	return &RPCServer{
		Impl: p.Impl,
	}, nil
}

func (p *RPCStatePlugin) Client(b *goplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{
		client: c,
	}, nil
}
