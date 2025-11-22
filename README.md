# Glide PHP Plugin

External PHP plugin for Glide - provides PHP and Composer integration.

## Overview

This plugin provides PHP functionality for Glide, including:

- Composer dependency management
- Framework detection (Laravel, Symfony, WordPress, Drupal, etc.)
- Testing tool detection and execution (PHPUnit, Pest, etc.)
- Static analysis tool integration (PHPStan, Psalm, Larastan, etc.)
- Project metadata extraction from composer.json

## Installation

### Method 1: Build from Source

```bash
# Clone the repository
git clone https://github.com/ivannovak/glide-plugin-php
cd glide-plugin-php

# Build the plugin
make build

# Install to PATH
sudo cp glide-plugin-php /usr/local/bin/
```

### Method 2: Go Install (when published)

```bash
go install github.com/ivannovak/glide-plugin-php/cmd/glide-plugin-php@latest
```

## Usage

Once installed, the plugin provides PHP commands to Glide:

```bash
# Install dependencies
glide install

# Install specific packages
glide install symfony/console

# Run Composer scripts
glide run test
glide run dev

# Run tests (auto-detects PHPUnit or Pest)
glide test
glide test --filter UserTest

# Run static analysis (auto-detects PHPStan, Psalm, or Larastan)
glide analyze
glide analyze src/
```

## Commands

### `install` (alias: `i`)

Install Composer dependencies or require new packages.

```bash
# Install all dependencies
glide install

# Require new packages
glide install symfony/console
glide install --dev phpunit/phpunit
```

### `run <script> [args...]`

Run any script defined in your `composer.json`:

```bash
glide run test
glide run dev
glide run post-install-cmd
```

### `test` (alias: `t`)

Run tests using the detected testing framework (PHPUnit or Pest).

```bash
# Run all tests
glide test

# Run specific tests
glide test --filter UserTest

# Run with coverage
glide test --coverage
```

### `analyze` (alias: `a`)

Run static analysis using detected tools (PHPStan, Psalm, or Larastan).

```bash
# Analyze entire project
glide analyze

# Analyze specific directory
glide analyze src/

# Run with specific level
glide analyze --level max
```

## Detection

The plugin automatically detects:

- **Frameworks**: Laravel, Symfony, WordPress, Drupal, Magento, CodeIgniter, Slim, Lumen, Laminas, Yii, CakePHP
- **Testing Tools**: PHPUnit, Pest, Codeception, Behat, PHPSpec
- **Quality Tools**: PHPStan, Psalm, PHP-CS-Fixer, PHPCS, PHPMD, Rector, Larastan
- **PHP Version**: From `require.php` in composer.json
- **Dependencies**: Installed status (vendor directory check)

## Development

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Tidy dependencies
make tidy
```

### Current Status

**Phase 3: PHP Plugin Creation** (In Progress)

The plugin structure is complete, but currently relies on Glide's internal packages via a local replace directive. To make this fully standalone:

1. Glide core needs to implement `sdk.RunPlugin()` function
2. SDK needs to expose extension data access in commands
3. Remove the local replace directive

This will be completed in a future phase when the public SDK API is finalized.

### Project Structure

```
glide-plugin-php/
├── cmd/
│   └── glide-plugin-php/
│       └── main.go              # Plugin entry point
├── internal/
│   ├── commands/
│   │   └── php.go               # PHP commands
│   └── plugin/
│       ├── detector.go          # PHP project detection
│       └── plugin.go            # Plugin implementation
├── Makefile                     # Build automation
└── README.md                    # This file
```

## License

MIT
