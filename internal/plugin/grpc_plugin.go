package plugin

import (
	v1 "github.com/ivannovak/glide/v2/pkg/plugin/sdk/v1"
)

// GRPCPlugin implements the gRPC GlidePluginServer interface
type GRPCPlugin struct {
	*v1.BasePlugin
}

// NewGRPCPlugin creates a new gRPC-based PHP plugin
func NewGRPCPlugin() *GRPCPlugin {
	metadata := &v1.PluginMetadata{
		Name:        "php",
		Version:     "1.0.0",
		Author:      "Glide Team",
		Description: "PHP framework detector for Glide",
		Homepage:    "https://github.com/ivannovak/glide-plugin-php",
		License:     "MIT",
		Tags:        []string{"language", "php", "composer", "laravel", "symfony", "detector"},
		Aliases:     []string{},
		Namespaced:  false,
	}

	p := &GRPCPlugin{
		BasePlugin: v1.NewBasePlugin(metadata),
	}

	// Note: This plugin only provides framework detection, not commands
	// Commands are handled by Glide's core CLI based on detected context

	return p
}

