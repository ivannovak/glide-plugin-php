package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ivannovak/glide/pkg/plugin/sdk"
	sdkv1 "github.com/ivannovak/glide/pkg/plugin/sdk/v1"
	"github.com/spf13/cobra"
)

// NewPHPCommands returns all PHP commands
func NewPHPCommands() []*sdk.PluginCommandDefinition {
	return []*sdk.PluginCommandDefinition{
		NewInstallCommand(),
		NewRunCommand(),
		NewTestCommand(),
		NewAnalyzeCommand(),
	}
}

// NewInstallCommand creates the 'install' command
func NewInstallCommand() *sdk.PluginCommandDefinition {
	return &sdk.PluginCommandDefinition{
		Name:  "install",
		Use:   "install [packages...]",
		Short: "Install PHP dependencies with Composer",
		Long: `Install PHP dependencies using Composer.

Without arguments, installs all dependencies from composer.json.
With arguments, requires the specified packages.

Examples:
  glide install                    # Install all dependencies
  glide install symfony/console    # Require a new package
  glide install --dev phpunit/phpunit  # Require a dev dependency
`,
		Aliases: []string{"i"},
		RunE:    executeInstall,
	}
}

// NewRunCommand creates the 'run' command
func NewRunCommand() *sdk.PluginCommandDefinition {
	return &sdk.PluginCommandDefinition{
		Name:  "run",
		Use:   "run <script> [args...]",
		Short: "Run a Composer script",
		Long: `Run any script defined in the composer.json scripts section.

Examples:
  glide run test               # Run the test script
  glide run dev                # Run the dev script
  glide run post-install-cmd   # Run a Composer script
`,
		Args: cobra.MinimumNArgs(1),
		RunE: executeRun,
	}
}

// NewTestCommand creates the 'test' command
func NewTestCommand() *sdk.PluginCommandDefinition {
	return &sdk.PluginCommandDefinition{
		Name:  "test",
		Use:   "test [args...]",
		Short: "Run PHPUnit or Pest tests",
		Long: `Run tests using the detected testing framework (PHPUnit or Pest).

Automatically detects which testing framework is installed and runs it.

Examples:
  glide test                   # Run all tests
  glide test --filter UserTest # Run specific test
  glide test --coverage        # Run with coverage
`,
		Aliases: []string{"t"},
		RunE:    executeTest,
	}
}

// NewAnalyzeCommand creates the 'analyze' command
func NewAnalyzeCommand() *sdk.PluginCommandDefinition {
	return &sdk.PluginCommandDefinition{
		Name:  "analyze",
		Use:   "analyze [paths...]",
		Short: "Run static analysis tools",
		Long: `Run static analysis using the detected tools (PHPStan, Psalm, etc.).

Automatically detects which analysis tools are installed and runs them.

Examples:
  glide analyze                # Analyze entire project
  glide analyze src/           # Analyze specific directory
  glide analyze --level max    # Run with maximum level
`,
		Aliases: []string{"a"},
		RunE:    executeAnalyze,
	}
}

// executeInstall runs the install command
func executeInstall(cmd *cobra.Command, args []string) error {
	// Get project context
	ctx := getProjectContext(cmd)
	if ctx == nil {
		return fmt.Errorf("project context not available")
	}

	// Build install command
	var installCmd *exec.Cmd
	if len(args) == 0 {
		// Install all dependencies
		installCmd = exec.Command("composer", "install")
		fmt.Println("Installing dependencies with Composer...")
	} else {
		// Require specific packages
		composerArgs := append([]string{"require"}, args...)
		installCmd = exec.Command("composer", composerArgs...)
		fmt.Printf("Installing packages: %v\n", args)
	}

	// Set working directory
	installCmd.Dir = ctx.Root
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	installCmd.Stdin = os.Stdin

	return installCmd.Run()
}

// executeRun runs a Composer script
func executeRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("script name required")
	}

	scriptName := args[0]
	scriptArgs := args[1:]

	// Get project context
	ctx := getProjectContext(cmd)
	if ctx == nil {
		return fmt.Errorf("project context not available")
	}

	// Build run command
	composerArgs := append([]string{"run-script", scriptName, "--"}, scriptArgs...)
	runCmd := exec.Command("composer", composerArgs...)

	// Set working directory
	runCmd.Dir = ctx.Root
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	runCmd.Stdin = os.Stdin

	fmt.Printf("Running script '%s' with Composer...\n", scriptName)
	return runCmd.Run()
}

// executeTest runs the detected test framework
func executeTest(cmd *cobra.Command, args []string) error {
	// Get project context
	ctx := getProjectContext(cmd)
	if ctx == nil {
		return fmt.Errorf("project context not available")
	}

	// Detect testing framework
	testingTool := detectTestingTool(ctx.Root)

	var testCmd *exec.Cmd
	switch testingTool {
	case "pest":
		testCmd = exec.Command("./vendor/bin/pest", args...)
		fmt.Println("Running Pest tests...")
	case "phpunit":
		testCmd = exec.Command("./vendor/bin/phpunit", args...)
		fmt.Println("Running PHPUnit tests...")
	default:
		return fmt.Errorf("no testing framework detected (PHPUnit or Pest)")
	}

	// Set working directory
	testCmd.Dir = ctx.Root
	testCmd.Stdout = os.Stdout
	testCmd.Stderr = os.Stderr
	testCmd.Stdin = os.Stdin

	return testCmd.Run()
}

// executeAnalyze runs static analysis tools
func executeAnalyze(cmd *cobra.Command, args []string) error {
	// Get project context
	ctx := getProjectContext(cmd)
	if ctx == nil {
		return fmt.Errorf("project context not available")
	}

	// Detect analysis tool
	analysisTool := detectAnalysisTool(ctx.Root)

	var analyzeCmd *exec.Cmd
	switch analysisTool {
	case "phpstan":
		cmdArgs := append([]string{"analyse"}, args...)
		analyzeCmd = exec.Command("./vendor/bin/phpstan", cmdArgs...)
		fmt.Println("Running PHPStan analysis...")
	case "psalm":
		analyzeCmd = exec.Command("./vendor/bin/psalm", args...)
		fmt.Println("Running Psalm analysis...")
	case "larastan":
		cmdArgs := append([]string{"analyse"}, args...)
		analyzeCmd = exec.Command("./vendor/bin/phpstan", cmdArgs...)
		fmt.Println("Running Larastan analysis...")
	default:
		return fmt.Errorf("no static analysis tool detected (PHPStan, Psalm, or Larastan)")
	}

	// Set working directory
	analyzeCmd.Dir = ctx.Root
	analyzeCmd.Stdout = os.Stdout
	analyzeCmd.Stderr = os.Stderr
	analyzeCmd.Stdin = os.Stdin

	return analyzeCmd.Run()
}

// getProjectContext extracts the project context from the command
func getProjectContext(cmd *cobra.Command) *sdkv1.ProjectContext {
	ctxValue := cmd.Context().Value("project_context")
	if ctxValue == nil {
		return nil
	}

	ctx, ok := ctxValue.(*sdkv1.ProjectContext)
	if !ok {
		return nil
	}

	return ctx
}

// detectTestingTool detects which testing framework is installed
func detectTestingTool(projectRoot string) string {
	// Check for Pest first (preferred for Laravel projects)
	if fileExists(projectRoot + "/vendor/bin/pest") {
		return "pest"
	}
	// Check for PHPUnit
	if fileExists(projectRoot + "/vendor/bin/phpunit") {
		return "phpunit"
	}
	return ""
}

// detectAnalysisTool detects which static analysis tool is installed
func detectAnalysisTool(projectRoot string) string {
	// Check for Larastan (PHPStan for Laravel)
	if fileExists(projectRoot + "/vendor/bin/phpstan") {
		// Check if larastan is in composer.json
		if fileExists(projectRoot + "/vendor/nunomaduro/larastan") {
			return "larastan"
		}
		return "phpstan"
	}
	// Check for Psalm
	if fileExists(projectRoot + "/vendor/bin/psalm") {
		return "psalm"
	}
	return ""
}

// fileExists checks if a file or directory exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
