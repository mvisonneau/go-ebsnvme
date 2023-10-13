# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.2.0] - 2023-10-13

### Changed

- CLI options implementation (no impact on the library though)

### Added

- Capability to output the device name instead of its full local path

## [v0.1.4] - 2023-10-13

### Changed

- Bumped to go 1.21
- Migrated CI from Drone to GH actions

## [0.1.3] - 2020-09-01

### Added

- Added gosec tests, excluding unsafe package warnings (G103)

### Changed

- Fixed a bug preventing device ids to be rendered correctly
- Bumped to go 1.15
- Bumped to gorelease 0.145.0
- Refactored codebase to follow golang standard structure

## [0.1.2] - 2020-09-01

### Changed

- Bumped to go 1.14
- Bumped to gorelease 0.142.0
- Fixed automated docker image builds
- Enhanced test and lint checks
- Refactored Makefile

## [0.1.1] - 2019-10-18

### Changed

- Manage dependencies using gomodules
- Upgraded to go 1.13
- Migrated CI to Drone
- Automated release of binaries, docker, DEB and RPM packages

## [0.1.0] - 2018-11-06

### Added

- Working state of the app and lib
- got some tests in place
- Makefile
- LICENSE
- README

[Unreleased]: https://github.com/mvisonneau/go-ebsnvme/compare/0.1.3...HEAD
[v0.2.0]: https://github.com/mvisonneau/go-ebsnvme/tree/0.1.3
[v0.1.4]: https://github.com/mvisonneau/go-ebsnvme/tree/0.1.3
[0.1.3]: https://github.com/mvisonneau/go-ebsnvme/tree/0.1.3
[0.1.2]: https://github.com/mvisonneau/go-ebsnvme/tree/0.1.2
[0.1.1]: https://github.com/mvisonneau/go-ebsnvme/tree/0.1.1
[0.1.0]: https://github.com/mvisonneau/go-ebsnvme/tree/0.1.0
