package plugin

import (
	"context"

	"github.com/ivannovak/glide-plugin-php/pkg/version"
	"github.com/ivannovak/glide/v3/pkg/plugin/sdk/v2"
)

// Config defines the plugin's type-safe configuration.
// Users configure this in .glide.yml under plugins.php
type Config struct {
	// PreferLaravel prioritizes Laravel detection over other frameworks
	PreferLaravel bool `json:"preferLaravel" yaml:"preferLaravel"`

	// PreferSymfony prioritizes Symfony detection over other frameworks
	PreferSymfony bool `json:"preferSymfony" yaml:"preferSymfony"`

	// EnableComposerScripts enables detection of composer scripts
	EnableComposerScripts bool `json:"enableComposerScripts" yaml:"enableComposerScripts"`
}

// DefaultConfig returns sensible defaults
func DefaultConfig() Config {
	return Config{
		PreferLaravel:         false,
		PreferSymfony:         false,
		EnableComposerScripts: true,
	}
}

// PHPPlugin implements the SDK v2 Plugin interface for PHP detection
type PHPPlugin struct {
	v2.BasePlugin[Config]
}

// New creates a new PHP plugin instance
func New() *PHPPlugin {
	return &PHPPlugin{}
}

// Metadata returns plugin information
func (p *PHPPlugin) Metadata() v2.Metadata {
	return v2.Metadata{
		Name:        "php",
		Version:     version.Version,
		Author:      "Glide Team",
		Description: "PHP framework detector for Glide",
		License:     "MIT",
		Homepage:    "https://github.com/ivannovak/glide-plugin-php",
		Tags:        []string{"language", "php", "composer", "laravel", "symfony", "detector"},
	}
}

// Configure is called with the type-safe configuration
func (p *PHPPlugin) Configure(ctx context.Context, config Config) error {
	return p.BasePlugin.Configure(ctx, config)
}

// Commands returns the list of commands this plugin provides.
// Note: This is a framework detector plugin, so it doesn't provide CLI commands.
func (p *PHPPlugin) Commands() []v2.Command {
	return []v2.Command{}
}

// Init is called once after plugin load
func (p *PHPPlugin) Init(ctx context.Context) error {
	return nil
}

// HealthCheck returns nil if the plugin is healthy
func (p *PHPPlugin) HealthCheck(ctx context.Context) error {
	return nil
}
