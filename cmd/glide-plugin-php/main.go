package main

import (
	"os"

	"github.com/ivannovak/glide-plugin-php/internal/plugin"
	sdk "github.com/ivannovak/glide/pkg/plugin/sdk/v1"
)

func main() {
	// Initialize the PHP gRPC plugin
	phpPlugin := plugin.NewGRPCPlugin()

	// Run the plugin using the SDK
	if err := sdk.RunPlugin(phpPlugin); err != nil {
		os.Exit(1)
	}
}
