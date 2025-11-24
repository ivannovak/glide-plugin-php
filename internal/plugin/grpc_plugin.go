package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
		Description: "PHP and Composer integration for Glide",
		Homepage:    "https://github.com/ivannovak/glide-plugin-php",
		License:     "MIT",
		Tags:        []string{"language", "php", "composer", "laravel", "symfony"},
		Aliases:     []string{},
		Namespaced:  false,
	}

	p := &GRPCPlugin{
		BasePlugin: v1.NewBasePlugin(metadata),
		
	}

	// Register all PHP commands
	p.registerCommands()

	return p
}

// registerCommands registers all PHP-related commands
func (p *GRPCPlugin) registerCommands() {
	// Install command
	p.RegisterCommand("install", v1.NewSimpleCommand(
		&v1.CommandInfo{
			Name:        "install",
			Description: "Install PHP dependencies with Composer",
			Category:    "dependencies",
			Aliases:     []string{"i"},
			Visibility:  "project-only",
		},
		p.executeInstall,
	))

	// Run command
	p.RegisterCommand("run", v1.NewSimpleCommand(
		&v1.CommandInfo{
			Name:        "run",
			Description: "Run a Composer script",
			Category:    "run",
			Visibility:  "project-only",
		},
		p.executeRun,
	))

	// Test command
	p.RegisterCommand("test", v1.NewSimpleCommand(
		&v1.CommandInfo{
			Name:        "test",
			Description: "Run PHPUnit or Pest tests",
			Category:    "test",
			Aliases:     []string{"t"},
			Visibility:  "project-only",
		},
		p.executeTest,
	))

	// Analyze command
	p.RegisterCommand("analyze", v1.NewSimpleCommand(
		&v1.CommandInfo{
			Name:        "analyze",
			Description: "Run static analysis tools",
			Category:    "lint",
			Aliases:     []string{"a"},
			Visibility:  "project-only",
		},
		p.executeAnalyze,
	))
}

// executeInstall runs the install command
func (p *GRPCPlugin) executeInstall(ctx context.Context, req *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
	workDir := req.WorkDir
	if workDir == "" {
		workDir = "."
	}

	var cmdParts []string
	if len(req.Args) == 0 {
		cmdParts = []string{"composer", "install"}
	} else {
		cmdParts = append([]string{"composer", "require"}, req.Args...)
	}

	return p.runCommand(ctx, cmdParts, workDir, req.Env)
}

// executeRun runs a Composer script
func (p *GRPCPlugin) executeRun(ctx context.Context, req *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
	if len(req.Args) == 0 {
		return &v1.ExecuteResponse{
			Success:  false,
			ExitCode: 1,
			Error:    "script name required",
		}, nil
	}

	workDir := req.WorkDir
	if workDir == "" {
		workDir = "."
	}

	cmdParts := append([]string{"composer", "run-script"}, req.Args...)
	return p.runCommand(ctx, cmdParts, workDir, req.Env)
}

// executeTest runs PHPUnit or Pest tests
func (p *GRPCPlugin) executeTest(ctx context.Context, req *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
	workDir := req.WorkDir
	if workDir == "" {
		workDir = "."
	}

	// Detect which test framework is available
	var cmdParts []string
	if p.hasVendorBin(workDir, "pest") {
		cmdParts = append([]string{"vendor/bin/pest"}, req.Args...)
	} else if p.hasVendorBin(workDir, "phpunit") {
		cmdParts = append([]string{"vendor/bin/phpunit"}, req.Args...)
	} else {
		return &v1.ExecuteResponse{
			Success:  false,
			ExitCode: 1,
			Error:    "No testing framework found (PHPUnit or Pest)",
		}, nil
	}

	return p.runCommand(ctx, cmdParts, workDir, req.Env)
}

// executeAnalyze runs static analysis tools
func (p *GRPCPlugin) executeAnalyze(ctx context.Context, req *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
	workDir := req.WorkDir
	if workDir == "" {
		workDir = "."
	}

	// Detect which analysis tool is available (prefer PHPStan, then Psalm, then Larastan)
	var cmdParts []string
	if p.hasVendorBin(workDir, "phpstan") {
		cmdParts = append([]string{"vendor/bin/phpstan", "analyze"}, req.Args...)
	} else if p.hasVendorBin(workDir, "psalm") {
		cmdParts = append([]string{"vendor/bin/psalm"}, req.Args...)
	} else if p.hasVendorBin(workDir, "larastan") {
		cmdParts = append([]string{"vendor/bin/phpstan", "analyze"}, req.Args...)
	} else {
		return &v1.ExecuteResponse{
			Success:  false,
			ExitCode: 1,
			Error:    "No static analysis tool found (PHPStan, Psalm, or Larastan)",
		}, nil
	}

	return p.runCommand(ctx, cmdParts, workDir, req.Env)
}

// hasVendorBin checks if a vendor binary exists
func (p *GRPCPlugin) hasVendorBin(workDir, binName string) bool {
	binPath := filepath.Join(workDir, "vendor", "bin", binName)
	_, err := os.Stat(binPath)
	return err == nil
}

// runCommand executes a command and returns the response
func (p *GRPCPlugin) runCommand(ctx context.Context, cmdParts []string, workDir string, env map[string]string) (*v1.ExecuteResponse, error) {
	if len(cmdParts) == 0 {
		return &v1.ExecuteResponse{
			Success:  false,
			ExitCode: 1,
			Error:    "empty command",
		}, nil
	}

	cmd := exec.CommandContext(ctx, cmdParts[0], cmdParts[1:]...)
	cmd.Dir = workDir

	// Set environment - start with parent environment
	cmd.Env = os.Environ()
	// Override/add custom environment variables
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	output, err := cmd.CombinedOutput()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return &v1.ExecuteResponse{
				Success:  false,
				ExitCode: 1,
				Error:    err.Error(),
			}, nil
		}
	}

	return &v1.ExecuteResponse{
		Success:  exitCode == 0,
		ExitCode: int32(exitCode),
		Stdout:   output,
	}, nil
}

