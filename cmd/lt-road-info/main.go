package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dimchansky/lt-road-info/internal/arcgis"
	"github.com/dimchansky/lt-road-info/internal/eismoinfo"
)

func main() {
	var (
		outputDir = flag.String("output", ".", "Output directory for GPX files")
		dataType  = flag.String("type", "all", "Type of data to download: all, restrictions, speed-control")
		verbose   = flag.Bool("verbose", false, "Enable verbose logging")
		help      = flag.Bool("help", false, "Show help message")
	)

	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	switch *dataType {
	case "all":
		downloadAll(*outputDir)
	case "restrictions":
		downloadRestrictions(*outputDir)
	case "speed-control":
		downloadSpeedControl(*outputDir)
	default:
		log.Fatalf("Unknown data type: %s. Use 'all', 'restrictions', or 'speed-control'", *dataType)
	}
}

func downloadAll(outputDir string) {
	downloadRestrictions(outputDir)
	downloadSpeedControl(outputDir)
}

func downloadRestrictions(outputDir string) {
	outputPath := filepath.Join(outputDir, "lt-road-restrictions.gpx")
	log.Printf("Downloading road restrictions to %s...", outputPath)

	if err := eismoinfo.DownloadRestrictions(outputPath); err != nil {
		log.Fatalf("Failed to download restrictions: %v", err)
	}

	log.Printf("Successfully downloaded road restrictions to %s", outputPath)
}

func downloadSpeedControl(outputDir string) {
	outputPath := filepath.Join(outputDir, "lt-speed-control.gpx")
	log.Printf("Downloading speed control sections to %s...", outputPath)

	if err := arcgis.DownloadSpeedControlSections(outputPath); err != nil {
		log.Fatalf("Failed to download speed control sections: %v", err)
	}

	log.Printf("Successfully downloaded speed control sections to %s", outputPath)
}

func printHelp() {
	fmt.Println("Lithuanian Road Information GPX Downloader")
	fmt.Println()
	fmt.Println("This tool downloads current road information from Lithuanian traffic systems")
	fmt.Println("and converts it to GPX format for use in navigation applications.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  lt-road-info [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  # Download all data to current directory")
	fmt.Println("  lt-road-info")
	fmt.Println()
	fmt.Println("  # Download only road restrictions")
	fmt.Println("  lt-road-info -type restrictions")
	fmt.Println()
	fmt.Println("  # Download to specific directory with verbose output")
	fmt.Println("  lt-road-info -output /path/to/gpx -verbose")
}
