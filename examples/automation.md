# Automation Examples

This guide shows how to automate downloading and processing of Lithuanian road information.

## Bash Script for Daily Updates

Create a script `update-road-info.sh`:

```bash
#!/bin/bash

# Configuration
OUTPUT_DIR="$HOME/Documents/GPX/Lithuanian-Roads"
LOG_FILE="$OUTPUT_DIR/update.log"

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Function to log messages
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$LOG_FILE"
}

# Download latest release
log "Starting download of Lithuanian road information..."

# Download road restrictions
curl -L -o "$OUTPUT_DIR/lt-road-restrictions.gpx" \
    "https://github.com/dimchansky/lt-road-info/releases/latest/download/lt-road-restrictions.gpx"

if [ $? -eq 0 ]; then
    log "Successfully downloaded road restrictions"
else
    log "Failed to download road restrictions"
fi

# Download speed control sections
curl -L -o "$OUTPUT_DIR/lt-speed-control.gpx" \
    "https://github.com/dimchansky/lt-road-info/releases/latest/download/lt-speed-control.gpx"

if [ $? -eq 0 ]; then
    log "Successfully downloaded speed control sections"
else
    log "Failed to download speed control sections"
fi

log "Update complete"
```

Make the script executable:
```bash
chmod +x update-road-info.sh
```

## Cron Job for Automatic Updates

Add to your crontab to run daily at 9 AM:

```bash
# Edit crontab
crontab -e

# Add this line:
0 9 * * * /path/to/update-road-info.sh
```

## Python Script with Notifications

```python
#!/usr/bin/env python3

import os
import requests
import logging
from datetime import datetime
from pathlib import Path

# Configuration
OUTPUT_DIR = Path.home() / "Documents" / "GPX" / "Lithuanian-Roads"
BASE_URL = "https://github.com/dimchansky/lt-road-info/releases/latest/download"
FILES = [
    "lt-road-restrictions.gpx",
    "lt-speed-control.gpx"
]

# Setup logging
OUTPUT_DIR.mkdir(parents=True, exist_ok=True)
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s',
    handlers=[
        logging.FileHandler(OUTPUT_DIR / "update.log"),
        logging.StreamHandler()
    ]
)

def download_file(filename):
    """Download a file from the latest release."""
    url = f"{BASE_URL}/{filename}"
    output_path = OUTPUT_DIR / filename
    
    try:
        response = requests.get(url, stream=True)
        response.raise_for_status()
        
        with open(output_path, 'wb') as f:
            for chunk in response.iter_content(chunk_size=8192):
                f.write(chunk)
        
        logging.info(f"Successfully downloaded {filename}")
        return True
    except Exception as e:
        logging.error(f"Failed to download {filename}: {e}")
        return False

def main():
    """Main function to download all GPX files."""
    logging.info("Starting Lithuanian road information update")
    
    success_count = 0
    for filename in FILES:
        if download_file(filename):
            success_count += 1
    
    if success_count == len(FILES):
        logging.info("All files downloaded successfully")
        # Add desktop notification (Linux example)
        os.system(f'notify-send "Road Info Updated" "Downloaded {success_count} GPX files"')
    else:
        logging.warning(f"Downloaded {success_count}/{len(FILES)} files")

if __name__ == "__main__":
    main()
```

## Docker Container for Scheduled Updates

Create a `Dockerfile`:

```dockerfile
FROM alpine:latest

RUN apk add --no-cache curl bash

COPY update-road-info.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/update-road-info.sh

# Run every day at 9 AM
RUN echo "0 9 * * * /usr/local/bin/update-road-info.sh" | crontab -

CMD ["crond", "-f"]
```

Build and run:
```bash
docker build -t lt-road-updater .
docker run -d -v ~/GPX:/output lt-road-updater
```