# Lithuanian Road Information GPX Downloader

[![Daily Update](https://github.com/dimchansky/lt-road-info/actions/workflows/update.yml/badge.svg)](https://github.com/dimchansky/lt-road-info/actions/workflows/update.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Automatically updated GPX files with Lithuanian road information including temporary restrictions (construction, repairs) and average speed control sections. Perfect for motorcyclists and drivers who want to import current road conditions into OsmAnd or other navigation apps.

## 📥 Download Latest GPX Files

The GPX files are automatically updated daily. You can download them directly:

- **[Road Restrictions (Latest)](https://github.com/dimchansky/lt-road-info/releases/latest/download/lt-road-restrictions.gpx)** - Temporary road restrictions, construction zones, repairs
- **[Speed Control Sections (Latest)](https://github.com/dimchansky/lt-road-info/releases/latest/download/lt-speed-control.gpx)** - Average speed control zones

## 🚀 Features

- **Daily Updates**: Automated GitHub Actions workflow fetches fresh data every day
- **Two Data Types**:
  - **Road Restrictions**: Construction sites, repairs, temporary closures, weight/height/width limits
  - **Speed Control Sections**: Average speed measurement zones
- **Navigation App Compatible**: GPX format works with OsmAnd, Garmin, and other navigation apps
- **Accurate Coordinates**: Proper conversion from Lithuanian LKS-94 to WGS-84 (GPS) coordinate system

## 📱 Using with Navigation Apps

### OsmAnd

1. Download the GPX files to your device
2. Copy files to `Android/data/net.osmand/files/tracks/` (or use OsmAnd's import feature)
3. Go to Configure Map → GPX files → select the imported files
4. The restrictions and speed zones will appear on your map

### Other Navigation Apps

Most navigation apps that support GPX import will work. Look for "Import GPX" or "Import Track" feature in your app.

## 🛠️ Building from Source

If you want to run the tool yourself or contribute to development:

### Prerequisites

- Go 1.21 or later
- Internet connection (for fetching data from Lithuanian traffic APIs)

### Installation

```bash
git clone https://github.com/dimchansky/lt-road-info.git
cd lt-road-info
go mod download
go build -o lt-road-info ./cmd/lt-road-info
```

### Usage

```bash
# Download all data (both restrictions and speed control)
./lt-road-info

# Download only road restrictions
./lt-road-info -type restrictions

# Download only speed control sections
./lt-road-info -type speed-control

# Specify output directory
./lt-road-info -output /path/to/gpx/files

# Enable verbose logging
./lt-road-info -verbose
```

### Command-line Options

- `-type` - Type of data to download: `all` (default), `restrictions`, `speed-control`
- `-output` - Output directory for GPX files (default: current directory)
- `-verbose` - Enable detailed logging
- `-help` - Show help message

### Development Tools

- `make test` - Run the test suite
- `make verify-coords` - Validate coordinate transformations with live data (see [cmd/verify-coords/README.md](cmd/verify-coords/README.md))

## 🔄 Data Sources

This tool fetches data from official Lithuanian road administration systems:

- **Road Restrictions**: [eismoinfo.lt](https://eismoinfo.lt) - Lithuanian Road Administration's traffic information portal
- **Speed Control Sections**: [gis.ktvis.lt](https://gis.ktvis.lt) - Lithuanian Transport Safety Administration's GIS system

## 📝 GPX File Structure

### Road Restrictions (`lt-road-restrictions.gpx`)

Each restriction is saved as a track with:
- **Name**: Type of restriction (e.g., "Road Works", "Road Closed", "Speed Limit")
- **Description**: Detailed information about the restriction
- **Track Points**: GPS coordinates forming the affected road section

### Speed Control Sections (`lt-speed-control.gpx`)

Each speed control section is saved as a track with:
- **Name**: Section identifier
- **Description**: Road name/number and speed limit (if available)
- **Track Points**: GPS coordinates of the speed measurement zone

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

This tool provides information from official sources but should not be your only source of road information. Always follow actual road signs and local regulations. The developers are not responsible for any issues arising from the use of this data.

## 🏍️ For Motorcyclists

This tool was created with motorcyclists in mind. Knowing about road restrictions and speed control zones in advance helps plan safer and more enjoyable routes. Stay safe and enjoy the ride!

---

Made with ❤️ for the Lithuanian riding community