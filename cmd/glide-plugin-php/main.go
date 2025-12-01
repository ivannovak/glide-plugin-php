package main

import (
	"fmt"
	"os"

	"github.com/ivannovak/glide-plugin-php/internal/plugin"
	"github.com/ivannovak/glide/v3/pkg/plugin/sdk/v2"
)

func main() {
	// Initialize the PHP plugin
	phpPlugin := plugin.New()

	// Run the plugin using SDK v2
	if err := v2.Serve(phpPlugin); err != nil {
		fmt.Fprintf(os.Stderr, "Plugin error: %v\n", err)
		os.Exit(1)
	}
}
