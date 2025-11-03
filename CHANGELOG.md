# Changelog

## [0.1.10](https://github.com/main-branch/chezroot/compare/v0.1.9...v0.1.10) (2025-11-03)


### Bug Fixes

* Add cache configuration for Go setup in CI workflows ([6092528](https://github.com/main-branch/chezroot/commit/6092528e5d23a14fb9cc78ed8baf4c373aabcaa4))


### Other Changes

* Remove notarization configuration from goreleaser ([9c6499d](https://github.com/main-branch/chezroot/commit/9c6499d038330168446699aa811331be984188b1))

## [0.1.9](https://github.com/main-branch/chezroot/compare/v0.1.8...v0.1.9) (2025-11-03)


### Features

* Publish chezroot to the main-branch/homebrew-tap ([17516a9](https://github.com/main-branch/chezroot/commit/17516a99c945c364a442282d6d9d9a5099871848))

## [0.1.8](https://github.com/main-branch/chezroot/compare/v0.1.7...v0.1.8) (2025-11-03)


### Other Changes

* Wait for chezroot notarization to complete before publishing ([83b2834](https://github.com/main-branch/chezroot/commit/83b2834fbd9bdf85b463dbf4dc2e9f5d11cf1484))

## [0.1.7](https://github.com/main-branch/chezroot/compare/v0.1.6...v0.1.7) (2025-11-03)


### Other Changes

* Add notarization configuration for macOS builds ([9c18828](https://github.com/main-branch/chezroot/commit/9c188288b2b76cbb2b7cf497bb30f86d47400f67))

## [0.1.6](https://github.com/main-branch/chezroot/compare/v0.1.5...v0.1.6) (2025-11-02)


### Other Changes

* Add a workflow to publish release artifacts with goreleaser ([f4f94ad](https://github.com/main-branch/chezroot/commit/f4f94ad42923ba04b47eabdc713e5dbda811c7fb))

## [0.1.5](https://github.com/main-branch/chezroot/compare/v0.1.4...v0.1.5) (2025-11-02)


### Other Changes

* Do not publish to homebrew tap ([cbe63b1](https://github.com/main-branch/chezroot/commit/cbe63b10f4e7c96f20a04a4272074e3124e6754a))

## [0.1.4](https://github.com/main-branch/chezroot/compare/v0.1.3...v0.1.4) (2025-11-01)


### Other Changes

* Cleanup release.yml ([08e39fa](https://github.com/main-branch/chezroot/commit/08e39faf65386743e67400f372937db447ed567b))

## [0.1.3](https://github.com/main-branch/chezroot/compare/v0.1.2...v0.1.3) (2025-11-01)


### Other Changes

* **release:** Implement decoupled Homebrew tap update ([1c50363](https://github.com/main-branch/chezroot/commit/1c5036304cf7e0177f3936b95fd44282be93775d))

## [0.1.2](https://github.com/main-branch/chezroot/compare/v0.1.1...v0.1.2) (2025-11-01)


### Other Changes

* Add homebrew-tap update to release-please config ([88d8700](https://github.com/main-branch/chezroot/commit/88d8700a85adbd46c8d4ce6492c7a6a46750f35c))
* Revert npm version to 0.0.1 and remove extra-files from release-please config ([fe8da7b](https://github.com/main-branch/chezroot/commit/fe8da7bfcffc07221d41044418e4bdaea7927f61))

## [0.1.1](https://github.com/main-branch/chezroot/compare/v0.1.0...v0.1.1) (2025-11-01)


### Bug Fixes

* Tell the general yaml linter not to lint the GitHub Actions files ([674a1f5](https://github.com/main-branch/chezroot/commit/674a1f56034c4559a8f78956df83b32a5ce005ff))


### Other Changes

* Add a Makefile and linters to the project ([38fb72e](https://github.com/main-branch/chezroot/commit/38fb72e8cdbfd03c5384b0b90e5d38e6df14ac81))
* Add ci build to Makefile ([b25a096](https://github.com/main-branch/chezroot/commit/b25a0968d73c1046a5d096945af8b37bdb8bb98e))
* Add dependabot integration on GitHub ([f60b102](https://github.com/main-branch/chezroot/commit/f60b102f6e53482079c0b884ba5071d3b8e6e34b))
* Add the continuous integration GitHub workflow ([1c2c3b3](https://github.com/main-branch/chezroot/commit/1c2c3b3730ed1f343ee0a32fa654b7a8d9afcda1))
* Change implementation from Bash to Go and update docs ([e90f2e0](https://github.com/main-branch/chezroot/commit/e90f2e0e63c014dd5cf58ccd1c2735675b483974))
* Create README.md to use as a specification for the chezroot tool ([6d1c5dd](https://github.com/main-branch/chezroot/commit/6d1c5dd6cad37af6b8e1a41ac35875102d3a992e))
* **deps-dev:** Bump @commitlint/cli from 19.8.1 to 20.1.0 ([039fe09](https://github.com/main-branch/chezroot/commit/039fe0903f89c69daf1e546eed5e87a12082b6e2))
* **deps-dev:** Bump @commitlint/config-conventional ([3cc9bb7](https://github.com/main-branch/chezroot/commit/3cc9bb75cb15224d606ea7c36673fedcfd50e387))
* **deps-dev:** Bump the npm-dev-deps group across 1 directory with 2 updates ([209560f](https://github.com/main-branch/chezroot/commit/209560feb0e3a7ccbc87a53b77169cb64bb5ecb1))
* Document the profile command and make other clarifications ([fea633a](https://github.com/main-branch/chezroot/commit/fea633a9c820bbd2526e452a20abadf0b905e3e9))
* Enforce conventional commits with a git commit-message hook ([7fb5f28](https://github.com/main-branch/chezroot/commit/7fb5f284dee950b33348e061a9807592ae2391dd))
* Enforce conventional commits with a GitHub workflow ([47f5b3a](https://github.com/main-branch/chezroot/commit/47f5b3a1dfc0de4ce439f1ceb8db4229393c3d0c))
* Exclude CHANGELOG.md from Markdown linting ([052e994](https://github.com/main-branch/chezroot/commit/052e9942dbc83b20220382bcd646ca7ea78e246b))
* Fix dependency type in the dependabot configuration ([769bdfa](https://github.com/main-branch/chezroot/commit/769bdfa485f157204183069a1bbb87d235cba612))
* Provide a trivial hello world implementation for chezroot ([dc3d91e](https://github.com/main-branch/chezroot/commit/dc3d91e76365730cb7b04cf8a86877c3e27721ca))
* **release:** Add release-please workflow and configuration ([0aa73c7](https://github.com/main-branch/chezroot/commit/0aa73c7defc2f25a1ecc062af08f079986d78be6))
* **release:** Update release-please config file paths in workflow ([351043e](https://github.com/main-branch/chezroot/commit/351043eb5d8610dd25e551897760c4475613d312))
* Update commit linting condition to skip dependabot PRs ([26ada20](https://github.com/main-branch/chezroot/commit/26ada20b91db428f5adeae7f80ba798902e7dbc1))
* Update design section of the README.md ([0cbf592](https://github.com/main-branch/chezroot/commit/0cbf5924c509e6ab5ce12339deeb60d991abbc36))
* Update Go and Node.js setup to use version files and improve dependency installation ([1bb610d](https://github.com/main-branch/chezroot/commit/1bb610d94f9d81f661aab57eef5d53efb40ec3ad))
