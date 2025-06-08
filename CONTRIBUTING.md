# Contributing to Lithuanian Road Information GPX Downloader

Thank you for your interest in contributing to this project! This guide will help you get started.

## Code of Conduct

This project follows a simple principle: be respectful and constructive in all interactions. We welcome contributions from everyone, regardless of experience level.

## Getting Started

### Development Setup

1. **Prerequisites**
   - Go 1.21 or later
   - Git
   - Internet connection (for testing with live APIs)

2. **Fork and Clone**
   ```bash
   git clone https://github.com/yourusername/lt-road-info.git
   cd lt-road-info
   ```

3. **Install Dependencies**
   ```bash
   go mod download
   ```

4. **Verify Setup**
   ```bash
   make test
   make verify-coords
   ```

## Development Guidelines

### Code Style

- **Formatting**: Run `gofmt -w .` before committing
- **Linting**: Code should pass `go vet` and `golangci-lint` (if available)
- **Documentation**: All exported functions and types must have Go doc comments
- **Error Handling**: Always handle errors appropriately, with context

### Testing Requirements

All contributions must include appropriate tests:

1. **Unit Tests**: For individual functions and methods
2. **Integration Tests**: For API interactions (use mocked responses when possible)
3. **Coordinate Validation**: Any changes to coordinate transformation must include validation tests

**Critical**: Any changes to coordinate transformation code MUST include tests that verify coordinates are in Lithuania (not Abu Dhabi or other wrong locations).

### Project Structure

```
lt-road-info/
├── cmd/                    # Command-line applications
├── internal/               # Private application code
│   ├── data/              # API clients and data types
│   ├── converter/         # GPX conversion logic
│   ├── transform/         # Coordinate transformation
│   ├── eismoinfo/         # Road restrictions API
│   └── arcgis/            # Speed control sections API
├── testdata/              # Test fixtures
└── examples/              # Usage examples and documentation
```

### Making Changes

1. **Create a Feature Branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make Your Changes**
   - Follow existing code patterns
   - Add tests for new functionality
   - Update documentation if needed

3. **Test Your Changes**
   ```bash
   make test-all
   make verify-coords
   ```

4. **Commit Your Changes**
   - Use clear, descriptive commit messages
   - Follow the pattern: `type: description`
   - Examples: `fix: correct coordinate transformation bug`, `feat: add timeout configuration`

## Types of Contributions

### Bug Reports

When reporting bugs, please include:
- Go version and operating system
- Steps to reproduce the issue
- Expected vs actual behavior
- Sample GPX files or coordinates if relevant

**For coordinate bugs**: Always include the problematic coordinates and expected geographic location.

### Feature Requests

Before implementing new features:
1. Open an issue to discuss the feature
2. Get feedback from maintainers
3. Consider backward compatibility
4. Ensure it aligns with the project goals

### Common Contribution Areas

- **API Enhancements**: Adding support for new Lithuanian road data sources
- **Output Formats**: Supporting additional GPS file formats
- **Testing**: Improving test coverage and reliability
- **Documentation**: Improving examples and usage guides
- **Performance**: Optimizing data processing and network requests

## Coordinate Transformation Guidelines

This project has specific requirements for coordinate handling:

### Critical Requirements

1. **Geographic Accuracy**: All coordinates must be in Lithuania (53.5°-56.5°N, 20.5°-27.0°E)
2. **Coordinate Order**: Always use `(latitude, longitude)` order consistently
3. **Transformation Testing**: Include tests that verify coordinates are geographically correct

### Before Submitting Coordinate Changes

Run the coordinate verification tool:
```bash
make verify-coords
```

This ensures that:
- Generated GPX files contain Lithuanian coordinates
- No coordinates appear in wrong countries (e.g., Abu Dhabi)
- Coordinate transformations are mathematically correct

## Pull Request Process

1. **Ensure CI Passes**: All GitHub Actions workflows must pass
2. **Update Documentation**: Update README.md if you change functionality
3. **Add Tests**: Include tests for new features or bug fixes
4. **Coordinate Verification**: For coordinate-related changes, demonstrate geographic accuracy

### Pull Request Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Performance improvement

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Coordinate verification passes (if applicable)
- [ ] Manual testing completed

## Geographic Validation (if applicable)
- [ ] Verified coordinates are in Lithuania
- [ ] No coordinates in Abu Dhabi or other wrong locations
- [ ] Coordinate transformation tests updated
```

## Release Process

The project uses automated releases:
- **Daily Updates**: GitHub Actions automatically generates fresh GPX files
- **Version Releases**: Created manually for significant changes
- **Latest Release**: Always contains the most recent GPX files

## Development Tools

### Available Make Targets

```bash
make test           # Run test suite
make test-all       # Run comprehensive tests including coordinate validation
make verify-coords  # Validate coordinates with live data
make build          # Build the binary
make fmt            # Format Go code
make lint           # Run linters (if available)
```

### Debugging Coordinate Issues

1. **Use Test Fixtures**: Create test cases with known LKS-94 coordinates
2. **Verify Transformations**: Use `cmd/verify-coords` to check real coordinates
3. **Check Boundaries**: Ensure all coordinates fall within Lithuanian boundaries
4. **Test Edge Cases**: Include coordinates from different regions of Lithuania

## Questions and Support

- **Issues**: Use GitHub Issues for bug reports and feature requests
- **Discussions**: Use GitHub Discussions for questions and ideas
- **Security**: See SECURITY.md for security-related reports

## License

By contributing to this project, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to make this tool better for the Lithuanian riding community! 🏍️🇱🇹