package plugin

import (
	"github.com/ivannovak/glide/pkg/plugin"
	"github.com/ivannovak/glide/pkg/plugin/sdk"
	"github.com/ivannovak/glide-plugin-php/internal/commands"
	"github.com/spf13/cobra"
)

// PHPPlugin implements the SDK Plugin interfaces for PHP functionality
type PHPPlugin struct {
	detector *PHPDetector
}

// New creates a new PHP plugin instance
func New() *PHPPlugin {
	return &PHPPlugin{
		detector: NewPHPDetector(),
	}
}

// Name returns the plugin identifier
func (p *PHPPlugin) Name() string {
	return "php"
}

// Version returns the plugin version
func (p *PHPPlugin) Version() string {
	return "1.0.0"
}

// Description returns the plugin description
func (p *PHPPlugin) Description() string {
	return "PHP and Composer integration for Glide"
}

// Register adds plugin commands to the command tree
func (p *PHPPlugin) Register(root *cobra.Command) error {
	// Get the command definitions from the SDK layer
	cmdDefs := p.ProvideCommands()

	// Convert and register each command with the root
	for _, cmdDef := range cmdDefs {
		if cmdDef != nil {
			cobraCmd := cmdDef.ToCobraCommand()

			// Wrap the command to inject project context
			// We need to do this because plugin commands don't have direct access to the app context
			p.wrapCommandWithContext(cobraCmd, root)

			root.AddCommand(cobraCmd)
		}
	}

	return nil
}

// wrapCommandWithContext wraps a command to inject project context from the root command
func (p *PHPPlugin) wrapCommandWithContext(cmd *cobra.Command, root *cobra.Command) {
	// Store the original RunE
	originalRunE := cmd.RunE
	if originalRunE == nil {
		return
	}

	// Wrap it to inject context
	cmd.RunE = func(c *cobra.Command, args []string) error {
		// Get the root command to access its context
		// The context should be set by the main CLI before execution
		rootCtx := c.Root().Context()
		if rootCtx != nil {
			// Set the context on this command
			c.SetContext(rootCtx)
		}

		// Call the original RunE
		return originalRunE(c, args)
	}
}

// Configure allows plugin-specific configuration
func (p *PHPPlugin) Configure(config map[string]interface{}) error {
	// PHP plugin doesn't require specific configuration yet
	// Future: Could add default Composer settings, PHP version preferences, etc.
	return nil
}

// Metadata returns plugin information
func (p *PHPPlugin) Metadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:        "php",
		Version:     "1.0.0",
		Author:      "Glide Team",
		Description: "PHP and Composer integration for Glide",
		Aliases:     []string{},
		Commands: []plugin.CommandInfo{
			{
				Name:        "install",
				Category:    "PHP",
				Description: "Install Composer dependencies",
				Aliases:     []string{"i"},
			},
			{
				Name:        "run",
				Category:    "PHP",
				Description: "Run Composer scripts",
				Aliases:     []string{},
			},
			{
				Name:        "test",
				Category:    "PHP",
				Description: "Run PHPUnit/Pest tests",
				Aliases:     []string{"t"},
			},
			{
				Name:        "analyze",
				Category:    "PHP",
				Description: "Run static analysis tools",
				Aliases:     []string{"a"},
			},
		},
		BuildTags:  []string{},
		ConfigKeys: []string{"php"},
	}
}

// ProvideContext returns the context extension for PHP detection
func (p *PHPPlugin) ProvideContext() sdk.ContextExtension {
	return p.detector
}

// ProvideCommands returns the commands provided by this plugin
func (p *PHPPlugin) ProvideCommands() []*sdk.PluginCommandDefinition {
	return commands.NewPHPCommands()
}
