package plugin_test

import (
	"fmt"
	"path"
	"testing"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/components-contrib/state/plugin"
	"github.com/dapr/kit/logger"
	"github.com/stretchr/testify/require"
)

func TestPluginRuns(t *testing.T) {
	store := plugin.NewPluginStateStore(logger.NewLogger("test"))
	metadataSet := []*state.Metadata{
		{
			Properties: map[string]string{
				plugin.MetadataName:          "go-memory",
				plugin.MetadataBaseDirectory: path.Join(".", "fixtures"),
				plugin.MetadataRunner:        string(plugin.RuntimeExec),
				plugin.MetadataVersion:       "0.0.1",
			},
		},
		{
			Properties: map[string]string{
				plugin.MetadataName:          "python-memory",
				plugin.MetadataBaseDirectory: path.Join(".", "fixtures"),
				plugin.MetadataRunner:        string(plugin.RuntimePython),
				plugin.MetadataVersion:       "0.0.1",
			},
		},
	}
	for _, metadata := range metadataSet {
		testName := fmt.Sprintf("can roundtrip %s", metadata.Properties[plugin.MetadataRunner])
		t.Run(testName, func(t *testing.T) {
			err := store.Init(*metadata)
			require.Nil(t, err)

			key := "hello"
			value := "world"

			etag := ""
			err = store.Set(&state.SetRequest{
				Key:      key,
				Value:    []byte(value),
				ETag:     &etag,
				Metadata: map[string]string{},
				Options:  state.SetStateOption{},
			})
			require.Nil(t, err)

			resp, err := store.Get(&state.GetRequest{
				Key:      key,
				Metadata: map[string]string{},
				Options:  state.GetStateOption{},
			})
			require.Nil(t, err)
			require.NotNil(t, resp)
			require.Equal(t, value, string(resp.Data))
		})
	}
}
