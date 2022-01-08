package plugin

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/dapr/components-contrib/state"
)

const (
	ComponentTypeState       ComponentType = "state"
	ComponentTypeSecretStore ComponentType = "secretstores"
	ComponentTypePubSub      ComponentType = "pubsub"
)

type PluginContext struct {
	Name          string
	BaseDirectory string
	Version       string
	Runner        string
	ComponentType ComponentType
}

func (pc *PluginContext) GeneratePluginFilePath(extension string) (string, error) {
	// <basedir>/<component_type/<name>/<version>/<file>
	relativePath := path.Join(
		pc.BaseDirectory,
		string(pc.ComponentType),
		pc.Name,
		pc.Version,
		pc.generatePluginFileName(extension),
	)
	return filepath.Abs(relativePath)
}

func (pc *PluginContext) generatePluginFileName(extension string) string {

	// dapr-store-<name>_v<version><extension>
	return fmt.Sprintf("dapr-%s-%s_v%s%s",
		pc.ComponentType,
		pc.Name,
		pc.Version,
		extension)
}

func PluginContextFromMetadata(
	metadata state.Metadata,
	componentType ComponentType) (*PluginContext, error) {

	result := &PluginContext{
		ComponentType: componentType,
	}
	var err error
	if result.BaseDirectory, err = parseMetadataString(metadata, MetadataBaseDirectory); err != nil {
		return nil, err
	}
	if result.Name, err = parseMetadataString(metadata, MetadataName); err != nil {
		return nil, err
	}
	if result.Runner, err = parseMetadataString(metadata, MetadataRunner); err != nil {
		return nil, err
	}
	if result.Version, err = parseMetadataString(metadata, MetadataVersion); err != nil {
		return nil, err
	}
	return result, nil
}

func parseMetadataString(metadata state.Metadata, key string) (string, error) {
	value, ok := metadata.Properties[key]
	if !ok {
		return "", fmt.Errorf("missing '%s' property", MetadataBaseDirectory)
	}
	return value, nil
}
