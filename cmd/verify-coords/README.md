# Coordinate Verification Tool

This tool validates that the Lithuanian road information GPX files contain geographically correct coordinates. It serves as a quality assurance tool to prevent coordinate transformation bugs.

## Purpose

The primary purpose is to catch coordinate transformation errors that could cause GPX tracks to appear in the wrong geographic location. For example, a lat/lon mixup bug could cause Lithuanian road tracks to appear in Abu Dhabi instead of Lithuania.

## What It Does

1. **Downloads Live Data**: Fetches current road restrictions and speed control data from official Lithuanian APIs
2. **Generates GPX Files**: Processes the data through the complete pipeline (parsing → coordinate transformation → GPX generation)
3. **Validates Geography**: Checks that all coordinates fall within Lithuanian boundaries
4. **Reports Statistics**: Shows the percentage of valid coordinates and provides sample coordinate values
5. **Detects Known Issues**: Specifically checks for coordinate mixup patterns that could place tracks in wrong countries

## Usage

### Command Line
```bash
# From project root
go run ./cmd/verify-coords

# Or using Make
make verify-coords
```

### Expected Output
```
🔍 Verifying coordinate transformations...

📍 Testing road restrictions...
   📊 restrictions: 10450/10450 coordinates in Lithuania (100.0%)
   📍 Sample coordinates: [[55.843685, 24.513895] [55.843720, 24.513943]]
✅ restrictions coordinates validation passed

🚗 Testing speed control sections...
   📊 speed control: 18301/18301 coordinates in Lithuania (100.0%)
   📍 Sample coordinates: [[54.651059, 25.426531] [54.651055, 25.426576]]
✅ speed control coordinates validation passed

✅ All coordinate transformations are correct!
🇱🇹 All GPX tracks are properly located in Lithuania
```

## Validation Criteria

### Geographic Boundaries
- **Lithuania Latitude**: 53.5° to 56.5° North
- **Lithuania Longitude**: 20.5° to 27.0° East

### Error Detection
- **Abu Dhabi Area**: 24.0-25.0°N, 54.0-56.0°E (indicates lat/lon mixup)
- **Outside Lithuania**: Any coordinates beyond Lithuanian boundaries
- **Minimum Threshold**: Requires >90% of coordinates to be in Lithuania

## When to Use

### Development
- **Before commits**: Verify coordinate transformations are working
- **After coordinate system changes**: Ensure transformations remain correct
- **Debugging geographic issues**: Get real coordinate samples for analysis

### Quality Assurance
- **Before releases**: Final validation of geographic accuracy
- **Manual testing**: Quick end-to-end verification
- **Troubleshooting**: When users report incorrect GPX locations

### Integration
- Could be integrated into CI/CD pipelines for automated geographic validation
- Useful for smoke testing in staging environments

## How It Works

1. **API Integration**: Uses the same data clients as the main application
2. **Real Data**: Downloads current data from live Lithuanian government APIs
3. **Complete Pipeline**: Tests the full data flow from API to GPX output
4. **Geographic Validation**: Applies boundary checks and coordinate range validation
5. **Regression Prevention**: Specifically checks for known error patterns

## Relationship to Other Tests

This tool complements the automated test suite:

- **Unit Tests**: Test individual functions with known test data
- **Integration Tests**: Test components with mocked API responses
- **Verify-Coords**: Tests the complete system with live data

It serves as the final verification step that ensures the application produces geographically accurate results in real-world conditions.

## Exit Codes

- **0**: All coordinates are valid (>90% in Lithuania)
- **1**: Validation failed or error occurred

This makes it suitable for use in scripts and automated validation pipelines.