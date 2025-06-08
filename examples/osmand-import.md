# Importing GPX files to OsmAnd

This guide shows how to import Lithuanian road information GPX files into OsmAnd navigation app.

## Method 1: Direct Import (Recommended)

1. Download GPX files from the [latest release](https://github.com/dimchansky/lt-road-info/releases/latest):
   - `lt-road-restrictions.gpx` - Road restrictions
   - `lt-speed-control.gpx` - Speed control sections

2. On your Android device, open the downloaded files with OsmAnd:
   - Navigate to your Downloads folder
   - Tap on each GPX file
   - Select "Open with OsmAnd"
   - Choose "Import as GPX track"

3. Enable the tracks on the map:
   - Open OsmAnd
   - Go to Menu → Configure map → GPX files
   - Enable the imported tracks

## Method 2: Manual Copy

1. Download the GPX files to your device

2. Copy files to OsmAnd tracks folder:
   - Connect your device to computer or use a file manager
   - Navigate to: `Android/data/net.osmand/files/tracks/`
   - Copy the GPX files to this folder

3. Restart OsmAnd and enable tracks as in Method 1

## Customizing Track Appearance

1. In Configure map → GPX files, tap on the track name
2. Select "Appearance"
3. Customize:
   - Color (recommended: Red for restrictions, Orange for speed control)
   - Width
   - Show arrows on track

## Tips

- Update the files regularly to get the latest road information
- You can create a "Road Info" subfolder in tracks to keep these files organized
- Consider setting different colors for different track types for easy identification