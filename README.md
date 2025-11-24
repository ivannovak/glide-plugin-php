# glide-plugin-php

[![CI](https://github.com/ivannovak/glide-plugin-php/actions/workflows/ci.yml/badge.svg)](https://github.com/ivannovak/glide-plugin-php/actions/workflows/ci.yml)
[![Semantic Release](https://github.com/ivannovak/glide-plugin-php/actions/workflows/semantic-release.yml/badge.svg)](https://github.com/ivannovak/glide-plugin-php/actions/workflows/semantic-release.yml)

PHP and Composer integration plugin for [Glide CLI](https://github.com/ivannovak/glide).

## Overview

This plugin provides PHP project detection and Composer integration for Glide. When installed, Glide will automatically detect PHP projects and provide intelligent commands for testing, static analysis, and dependency management.

## Installation

### From GitHub Releases (Recommended)

```bash
glide plugins install github.com/ivannovak/glide-plugin-php
```

### From Source

```bash
# Clone the repository
git clone https://github.com/ivannovak/glide-plugin-php.git
cd glide-plugin-php

# Build and install (requires Go 1.24+)
make install
```

## What It Detects

The plugin automatically detects PHP projects by looking for:

- **Required files**: `composer.json`
- **Lock files**: `composer.lock`
- **Directories**: `vendor/`
- **Config files**: `phpunit.xml`, `phpstan.neon`, `psalm.xml`

### Framework Detection

The plugin recognizes popular PHP frameworks:

- **Laravel** - Full-stack framework
- **Symfony** - Enterprise framework
- **WordPress** - CMS platform
- **Drupal** - CMS platform
- **Magento** - E-commerce platform
- **CodeIgniter** - Lightweight framework
- **Slim** - Micro framework
- **Lumen** - Laravel micro-framework
- **Laminas** (formerly Zend)
- **Yii** - High-performance framework
- **CakePHP** - Rapid development framework

### Tool Detection

The plugin automatically detects and integrates with:

**Testing Tools:**
- PHPUnit
- Pest
- Codeception
- Behat
- PHPSpec

**Quality Tools:**
- PHPStan
- Psalm
- Larastan (Laravel-specific PHPStan)
- PHP-CS-Fixer
- PHPCS (PHP CodeSniffer)
- PHPMD (PHP Mess Detector)
- Rector

## Available Commands

Once a PHP project is detected, the following commands become available:

### Dependency Management
- `install` (alias: `i`) - Install Composer dependencies
  - With args: Requires specified packages (`composer require`)
  - Without args: Installs all dependencies (`composer install`)

### Script Execution
- `run <script>` - Run a composer.json script
  - Executes scripts defined in composer.json

### Testing
- `test` (alias: `t`) - Run tests with auto-detected framework
  - Automatically uses PHPUnit or Pest
  - Forwards all arguments to the testing tool

### Static Analysis
- `analyze` (alias: `a`) - Run static analysis
  - Auto-detects PHPStan, Psalm, or Larastan
  - Forwards all arguments to the analysis tool

## Configuration

The plugin works out-of-the-box without configuration. However, you can customize behavior in your `.glide.yml`:

```yaml
plugins:
  php:
    enabled: true
    # Additional configuration options can be added here in the future
```

## Examples

### Basic PHP Project

```bash
# Navigate to your PHP project
cd my-laravel-app

# Glide automatically detects PHP and Laravel
glide help

# Install dependencies
glide install

# Install a specific package
glide install symfony/console

# Run tests
glide test

# Run static analysis
glide analyze
```

### Testing Workflows

```bash
# Run all tests
glide test

# Run specific test file
glide test tests/Unit/UserTest.php

# Run with filter
glide test --filter=UserTest

# Run with coverage
glide test --coverage
```

### Static Analysis

```bash
# Analyze entire project
glide analyze

# Analyze specific directory
glide analyze src/

# Run with specific level (PHPStan)
glide analyze --level=max

# Run with configuration
glide analyze --configuration=phpstan.neon
```

### Composer Scripts

```bash
# Run any composer script
glide run test
glide run dev
glide run post-install-cmd
glide run lint
```

## Development

### Prerequisites

- Go 1.24 or higher
- Make (optional, for convenience targets)

### Building

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage

# Run linters
make lint

# Format code
make fmt
```

### Testing

The plugin includes comprehensive tests for:

- PHP version detection
- Framework detection
- Tool detection (testing & quality)
- Composer metadata extraction
- Command execution

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass (`make test`)
6. Submit a pull request

## License

MIT License - see [LICENSE](LICENSE) for details.

## Related Projects

- [Glide](https://github.com/ivannovak/glide) - The main Glide CLI
- [glide-plugin-go](https://github.com/ivannovak/glide-plugin-go) - Go plugin for Glide
- [glide-plugin-node](https://github.com/ivannovak/glide-plugin-node) - Node.js plugin for Glide
- [glide-plugin-docker](https://github.com/ivannovak/glide-plugin-docker) - Docker plugin for Glide

## Support

- [GitHub Issues](https://github.com/ivannovak/glide-plugin-php/issues)
- [Glide Documentation](https://github.com/ivannovak/glide#readme)
- [Plugin Development Guide](https://github.com/ivannovak/glide/blob/main/docs/PLUGIN_DEVELOPMENT.md)
