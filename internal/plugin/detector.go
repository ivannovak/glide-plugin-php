package plugin

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// PHPDetector implements the SDK ContextExtension interface for PHP detection
type PHPDetector struct{}

// ComposerJSON represents the structure of composer.json
type ComposerJSON struct {
	Name              string                       `json:"name"`
	Description       string                       `json:"description"`
	Version           string                       `json:"version"`
	Type              string                       `json:"type"`
	License           interface{}                  `json:"license"` // Can be string or []string
	Require           map[string]string            `json:"require"`
	RequireDev        map[string]string            `json:"require-dev"`
	Autoload          map[string]interface{}       `json:"autoload"`
	AutoloadDev       map[string]interface{}       `json:"autoload-dev"`
	Scripts           map[string]interface{}       `json:"scripts"`
	Config            map[string]interface{}       `json:"config"`
	Extra             map[string]interface{}       `json:"extra"`
	Minimum           map[string]interface{}       `json:"minimum-stability"`
	PreferStable      bool                         `json:"prefer-stable"`
}

// NewPHPDetector creates a new PHP detector
func NewPHPDetector() *PHPDetector {
	return &PHPDetector{}
}

// Name returns the unique identifier for this extension
func (d *PHPDetector) Name() string {
	return "php"
}

// Detect analyzes the project environment and returns PHP-specific context data
func (d *PHPDetector) Detect(ctx context.Context, projectRoot string) (interface{}, error) {
	// Check if composer.json exists
	composerJSONPath := filepath.Join(projectRoot, "composer.json")
	if _, err := os.Stat(composerJSONPath); os.IsNotExist(err) {
		// No composer.json, not a PHP project
		return nil, nil
	}

	// Read and parse composer.json
	composer, err := d.readComposerJSON(composerJSONPath)
	if err != nil {
		// composer.json exists but can't be read/parsed
		return map[string]interface{}{
			"php_detected": true,
			"error":        "failed to parse composer.json",
		}, nil
	}

	// Detect frameworks
	frameworks := d.detectFrameworks(composer, projectRoot)

	// Detect testing tools
	testingTools := d.detectTestingTools(composer)

	// Detect quality tools
	qualityTools := d.detectQualityTools(composer)

	// Build the extension data structure
	result := map[string]interface{}{
		"php_detected": true,
		"project_name": composer.Name,
	}

	// Add optional fields
	if composer.Description != "" {
		result["description"] = composer.Description
	}
	if composer.Version != "" {
		result["version"] = composer.Version
	}
	if composer.Type != "" {
		result["project_type"] = composer.Type
	}
	if phpVersion, ok := composer.Require["php"]; ok {
		result["php_version"] = phpVersion
	}
	if len(frameworks) > 0 {
		result["frameworks"] = frameworks
	}
	if len(testingTools) > 0 {
		result["testing_tools"] = testingTools
	}
	if len(qualityTools) > 0 {
		result["quality_tools"] = qualityTools
	}

	// Check for vendor directory
	vendorPath := filepath.Join(projectRoot, "vendor")
	if stat, err := os.Stat(vendorPath); err == nil && stat.IsDir() {
		result["dependencies_installed"] = true
	}

	return result, nil
}

// Merge merges two context data structures
func (d *PHPDetector) Merge(existing, new interface{}) (interface{}, error) {
	// If either is nil, return the non-nil one
	if existing == nil {
		return new, nil
	}
	if new == nil {
		return existing, nil
	}

	// Type assert both to maps
	existingMap, ok1 := existing.(map[string]interface{})
	newMap, ok2 := new.(map[string]interface{})

	if !ok1 || !ok2 {
		// If either is not a map, prefer new
		return new, nil
	}

	// Merge maps (new values override existing)
	result := make(map[string]interface{})
	for k, v := range existingMap {
		result[k] = v
	}
	for k, v := range newMap {
		result[k] = v
	}

	return result, nil
}

// readComposerJSON reads and parses composer.json
func (d *PHPDetector) readComposerJSON(path string) (*ComposerJSON, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var composer ComposerJSON
	if err := json.Unmarshal(data, &composer); err != nil {
		return nil, err
	}

	return &composer, nil
}

// detectFrameworks detects PHP frameworks from composer.json and file existence
func (d *PHPDetector) detectFrameworks(composer *ComposerJSON, projectRoot string) []string {
	var frameworks []string

	// Combine all dependencies
	allDeps := make(map[string]bool)
	for dep := range composer.Require {
		allDeps[dep] = true
	}
	for dep := range composer.RequireDev {
		allDeps[dep] = true
	}

	// Check for frameworks by package name
	frameworkChecks := map[string]string{
		"laravel/framework":            "Laravel",
		"symfony/symfony":              "Symfony",
		"symfony/framework-bundle":     "Symfony",
		"cakephp/cakephp":              "CakePHP",
		"yiisoft/yii2":                 "Yii",
		"codeigniter4/framework":       "CodeIgniter",
		"slim/slim":                    "Slim",
		"laravel/lumen-framework":      "Lumen",
		"laminas/laminas-mvc":          "Laminas",
		"magento/product-community-edition": "Magento",
	}

	for dep, name := range frameworkChecks {
		if allDeps[dep] {
			frameworks = append(frameworks, name)
		}
	}

	// Check for WordPress (wp-config.php)
	wpConfigPath := filepath.Join(projectRoot, "wp-config.php")
	if _, err := os.Stat(wpConfigPath); err == nil {
		frameworks = append(frameworks, "WordPress")
	}

	// Check for Drupal (core/lib/Drupal.php)
	drupalPath := filepath.Join(projectRoot, "core", "lib", "Drupal.php")
	if _, err := os.Stat(drupalPath); err == nil {
		frameworks = append(frameworks, "Drupal")
	}

	return frameworks
}

// detectTestingTools detects testing frameworks and tools
func (d *PHPDetector) detectTestingTools(composer *ComposerJSON) []string {
	var tools []string

	// Combine all dependencies (testing tools are usually in require-dev)
	allDeps := make(map[string]bool)
	for dep := range composer.Require {
		allDeps[dep] = true
	}
	for dep := range composer.RequireDev {
		allDeps[dep] = true
	}

	// Check for testing tools
	testingChecks := map[string]string{
		"phpunit/phpunit":     "PHPUnit",
		"pestphp/pest":        "Pest",
		"codeception/codeception": "Codeception",
		"behat/behat":         "Behat",
		"phpspec/phpspec":     "PHPSpec",
	}

	for dep, name := range testingChecks {
		if allDeps[dep] {
			tools = append(tools, name)
		}
	}

	return tools
}

// detectQualityTools detects code quality and static analysis tools
func (d *PHPDetector) detectQualityTools(composer *ComposerJSON) []string {
	var tools []string

	// Combine all dependencies
	allDeps := make(map[string]bool)
	for dep := range composer.Require {
		allDeps[dep] = true
	}
	for dep := range composer.RequireDev {
		allDeps[dep] = true
	}

	// Check for quality tools
	qualityChecks := map[string]string{
		"phpstan/phpstan":                    "PHPStan",
		"vimeo/psalm":                        "Psalm",
		"friendsofphp/php-cs-fixer":          "PHP-CS-Fixer",
		"squizlabs/php_codesniffer":          "PHPCS",
		"phpmd/phpmd":                        "PHPMD",
		"rector/rector":                      "Rector",
		"larastan/larastan":                  "Larastan",
	}

	for dep, name := range qualityChecks {
		if allDeps[dep] || d.hasPartialMatch(allDeps, dep) {
			tools = append(tools, name)
		}
	}

	return tools
}

// hasPartialMatch checks if any dependency starts with the given prefix
func (d *PHPDetector) hasPartialMatch(deps map[string]bool, prefix string) bool {
	for dep := range deps {
		if strings.HasPrefix(dep, prefix) {
			return true
		}
	}
	return false
}
