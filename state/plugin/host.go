package plugin

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/components-contrib/state/plugin/shared"
	"github.com/dapr/kit/logger"
	"github.com/hashicorp/go-plugin"
)

// NewPluginStateStore creates a new instance of a Sql Server transaction store.
func NewPluginStateStore(logger logger.Logger) state.Store {
	return &Host{
		logger: logger,
	}
}

type Host struct {
	store  state.Store
	client plugin.Client
	logger logger.Logger
}

func (h *Host) Features() []state.Feature {
	return []state.Feature{}
}

const (
	MetadataPrefix        = "plugin."
	MetadataBaseDirectory = MetadataPrefix + "basedir"
	MetadataVersion       = MetadataPrefix + "version"
	MetadataName          = MetadataPrefix + "name"
	MetadataRunner        = MetadataPrefix + "runner"
)

func (h *Host) Init(metadata state.Metadata) error {
	pluginContext, err := PluginContextFromMetadata(metadata, ComponentTypeState)
	if err != nil {
		return err
	}
	runtimeContext := GetRuntimeContext(Runtime(pluginContext.Runner))
	path, err := pluginContext.GeneratePluginFilePath(runtimeContext.Extension())
	if err != nil {
		return err
	}

	// We're a host. Start by launching the plugin process.
	client := h.createClient(runtimeContext.Command(path))

	// connect via rpc
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	// request the plugin
	raw, err := rpcClient.Dispense(shared.ProtocolGRPC)
	if err != nil {
		return err
	}

	// cast to a store
	var ok bool
	h.store, ok = raw.(state.Store)
	if !ok {
		return fmt.Errorf("expected type %T to be state.Store", raw)
	}

	// call the plugin, filter out any host metadata that isn't needed by the downstream
	return h.store.Init(h.filterMetadata(metadata))
}

func (h *Host) createClient(cmd *exec.Cmd) *plugin.Client {
	return plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             cmd,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC,
			plugin.ProtocolGRPC,
		},
	})
}

// filterMetadata removes any "plugin." prefixed parameters
func (h *Host) filterMetadata(metadata state.Metadata) state.Metadata {
	properties := map[string]string{}
	for k, v := range metadata.Properties {
		if strings.HasPrefix(k, MetadataPrefix) {
			continue
		}
		properties[k] = v
	}
	return state.Metadata{
		Properties: properties,
	}
}

func (h *Host) Get(req *state.GetRequest) (*state.GetResponse, error) {
	return h.store.Get(req)
}

func (h *Host) Set(req *state.SetRequest) error {
	return h.store.Set(req)
}

func (h *Host) Ping() error {
	return h.store.Ping()
}

func (h *Host) Delete(req *state.DeleteRequest) error {
	return h.store.Delete(req)
}

func (h *Host) BulkDelete(req []state.DeleteRequest) error {
	return h.store.BulkDelete(req)
}

func (h *Host) BulkGet(req []state.GetRequest) (bool, []state.BulkGetResponse, error) {
	return h.store.BulkGet(req)
}

func (h *Host) BulkSet(req []state.SetRequest) error {
	return h.store.BulkSet(req)
}

func (h *Host) Close() error {
	h.client.Kill()
	return nil
}
