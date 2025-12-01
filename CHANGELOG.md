## [3.0.0](https://github.com/ivannovak/glide-plugin-php/compare/v2.3.3...v3.0.0) (2025-12-01)


### ⚠ BREAKING CHANGES

* This plugin now requires Glide v2.4.0+ with SDK v2.

Migration to SDK v2:
- Replace v1.BasePlugin with v2.BasePlugin[Config]
- Add type-safe Config struct with preferLaravel, preferSymfony,
  and enableComposerScripts options
- Update main.go to use v2.Serve()
- Create new plugin.go with SDK v2 patterns
- Remove legacy grpc_plugin.go (v1 implementation)

The plugin now uses the declarative SDK v2 pattern with:
- Type-safe configuration via Go generics
- Unified lifecycle management
- Declarative metadata via Metadata() method

Note: go.mod includes a replace directive pointing to local glide
repo until v2.4.0 with SDK v2 is released.

* feat!: upgrade to glide SDK v3.0.0
* Updates module dependency from glide/v2 to glide/v3 v3.0.0.
This aligns with the SDK v2 type-safe configuration system released in glide v3.0.0.

- Update go.mod to require github.com/ivannovak/glide/v3 v3.0.0
- Remove local replace directive (now using published version)
- Update all imports from /v2/ to /v3/

* ci: add CI workflow for PR validation

### Features

* upgrade to glide SDK v3.0.0 ([#1](https://github.com/ivannovak/glide-plugin-php/issues/1)) ([c801202](https://github.com/ivannovak/glide-plugin-php/commit/c8012021802a2717554749a5ec4e46fd8cbde1f1))

## [2.3.3](https://github.com/ivannovak/glide-plugin-php/compare/v2.3.2...v2.3.3) (2025-11-25)


### Bug Fixes

* remove command registration from detector-only plugin ([6d1cbd5](https://github.com/ivannovak/glide-plugin-php/commit/6d1cbd5b067c5d0e0ec104a7414997998479d468))

## [2.3.2](https://github.com/ivannovak/glide-plugin-php/compare/v2.3.1...v2.3.2) (2025-11-25)


### Bug Fixes

* correct build path in release workflow ([6393da7](https://github.com/ivannovak/glide-plugin-php/commit/6393da73a909d02587e44a0f2424970b56eb4549))

## [2.3.1](https://github.com/ivannovak/glide-plugin-php/compare/v2.3.0...v2.3.1) (2025-11-25)


### Bug Fixes

* remove CI dependency from release workflow ([8b4248b](https://github.com/ivannovak/glide-plugin-php/commit/8b4248b92ae99a715602d0d8b356e43283fe8fca))

## [2.3.0](https://github.com/ivannovak/glide-plugin-php/compare/v2.2.0...v2.3.0) (2025-11-25)


### Features

* use published Glide v2.2.0 ([7afeaec](https://github.com/ivannovak/glide-plugin-php/commit/7afeaeccc00fbd439d0ec604a508e11f5a9d45c5))

## [2.2.0](https://github.com/ivannovak/glide-plugin-php/compare/v2.1.0...v2.2.0) (2025-11-24)


### Features

* migrate to Glide v2 module path ([b17605c](https://github.com/ivannovak/glide-plugin-php/commit/b17605ce356e7655a1e989ada7278c4a292e1789))

## [2.1.0](https://github.com/ivannovak/glide-plugin-php/compare/v2.0.0...v2.1.0) (2025-11-24)


### Features

* add release workflow for cross-platform binaries ([ec301be](https://github.com/ivannovak/glide-plugin-php/commit/ec301be13dbbff70d367335c0df1a7b31ef4f9ba))

## [2.0.0](https://github.com/ivannovak/glide-plugin-php/compare/v1.0.0...v2.0.0) (2025-11-24)


### ⚠ BREAKING CHANGES

* Plugin now uses gRPC instead of library architecture

### Bug Fixes

* **build:** correct module name to glide-plugin-php ([fbdade3](https://github.com/ivannovak/glide-plugin-php/commit/fbdade3d8f894a1cfce636b9380d5418a479b619))


### Code Refactoring

* migrate to gRPC architecture and cleanup legacy code ([21b9f58](https://github.com/ivannovak/glide-plugin-php/commit/21b9f584eaf5ba536a1af9a26e558c3c8791c153))

## 1.0.0 (2025-11-22)


### Features

* initial PHP plugin implementation ([b76a967](https://github.com/ivannovak/glide-plugin-php/commit/b76a967e3528959042d35d65a1ed370fadc4b2a5))

## 1.0.0 (2025-11-21)


### Features

* initial PHP plugin implementation ([41e4220](https://github.com/ivannovak/glide-plugin-php/commit/41e42206488f4ef497ab1fada18a171aae44dd1a))


### Bug Fixes

* update .gitignore to allow cmd/glide-plugin-php directory ([a4b6aca](https://github.com/ivannovak/glide-plugin-php/commit/a4b6aca21c240e6c6bdd279ccfd1eeebd2a24ca9))
* update package.json repository URL to glide-plugin-php ([09d552b](https://github.com/ivannovak/glide-plugin-php/commit/09d552b628294f7c3572f25e4bcda97b3db432b7))

## 1.0.0 (2025-11-21)


### Features

* **plugin:** initial Docker plugin extraction (Phase 6) ([545ae53](https://github.com/ivannovak/glide-plugin-docker/commit/545ae5308df59fbc0e446339fbafbce719b74892))
