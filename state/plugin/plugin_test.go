package plugin_test

import (
	"path"
	"testing"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/components-contrib/state/plugin"
	"github.com/dapr/kit/logger"
	"github.com/stretchr/testify/require"
)

func TestGoPluginRun(t *testing.T) {
	store := plugin.NewPluginStateStore(logger.NewLogger("test"))
	t.Run("can roundtrip", func(t *testing.T) {
		metadata := state.Metadata{
			Properties: map[string]string{
				plugin.MetadataName:          "go-memory",
				plugin.MetadataBaseDirectory: path.Join(".", "fixtures"),
				plugin.MetadataRunner:        string(plugin.RuntimeExec),
				plugin.MetadataVersion:       "0.0.1",
			},
		}
		err := store.Init(metadata)
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
