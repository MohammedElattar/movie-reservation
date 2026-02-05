#!/bin/bash
# update-go.sh - Automatically update Go to the latest version

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Go Version Updater ===${NC}\n"

# Get current Go version
CURRENT_VERSION=$(go version 2>/dev/null | awk '{print $3}' | sed 's/go//' || echo "not installed")
echo -e "Current Go version: ${YELLOW}${CURRENT_VERSION}${NC}"

# Detect OS and Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture names
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    armv6l)
        ARCH="armv6l"
        ;;
    i686|i386)
        ARCH="386"
        ;;
esac

echo -e "Detected OS: ${YELLOW}${OS}${NC}, Architecture: ${YELLOW}${ARCH}${NC}\n"

# Fetch latest Go version from official website
echo "Fetching latest Go version..."
LATEST_VERSION=$(curl -s https://go.dev/VERSION?m=text | head -n 1 | sed 's/go//')

if [ -z "$LATEST_VERSION" ]; then
    echo -e "${RED}Error: Could not fetch latest Go version${NC}"
    exit 1
fi

echo -e "Latest Go version: ${GREEN}${LATEST_VERSION}${NC}\n"

# Check if already up to date
if [ "$CURRENT_VERSION" = "$LATEST_VERSION" ]; then
    echo -e "${GREEN}✓ Go is already up to date!${NC}"
    exit 0
fi

# Ask for confirmation
echo -e "${YELLOW}Update from ${CURRENT_VERSION} to ${LATEST_VERSION}?${NC}"
read -p "Continue? (y/n) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Update cancelled."
    exit 0
fi

# Determine download URL
DOWNLOAD_URL="https://go.dev/dl/go${LATEST_VERSION}.${OS}-${ARCH}.tar.gz"
DOWNLOAD_FILE="/tmp/go${LATEST_VERSION}.${OS}-${ARCH}.tar.gz"

echo -e "\n${BLUE}Downloading Go ${LATEST_VERSION}...${NC}"
if ! curl -L --progress-bar "$DOWNLOAD_URL" -o "$DOWNLOAD_FILE"; then
    echo -e "${RED}Error: Download failed${NC}"
    exit 1
fi

# Verify download
if [ ! -f "$DOWNLOAD_FILE" ]; then
    echo -e "${RED}Error: Download file not found${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Download complete${NC}\n"

# Backup and remove old Go installation
echo -e "${BLUE}Removing old Go installation...${NC}"

# Detect Go installation path
if [ -d "/usr/local/go" ]; then
    INSTALL_PATH="/usr/local/go"
elif [ -d "$HOME/.go" ]; then
    INSTALL_PATH="$HOME/.go"
elif [ -d "/opt/go" ]; then
    INSTALL_PATH="/opt/go"
else
    INSTALL_PATH="/usr/local/go"
fi

# Check if we need sudo
NEED_SUDO=false
if [[ "$INSTALL_PATH" == /usr/* ]] || [[ "$INSTALL_PATH" == /opt/* ]]; then
    NEED_SUDO=true
fi

# Remove old installation
if [ -d "$INSTALL_PATH" ]; then
    if [ "$NEED_SUDO" = true ]; then
        echo "Removing $INSTALL_PATH (requires sudo)..."
        sudo rm -rf "$INSTALL_PATH"
    else
        echo "Removing $INSTALL_PATH..."
        rm -rf "$INSTALL_PATH"
    fi
fi

# Extract new Go version
echo -e "${BLUE}Installing Go ${LATEST_VERSION}...${NC}"

TARGET_DIR=$(dirname "$INSTALL_PATH")

if [ "$NEED_SUDO" = true ]; then
    sudo tar -C "$TARGET_DIR" -xzf "$DOWNLOAD_FILE"
else
    mkdir -p "$TARGET_DIR"
    tar -C "$TARGET_DIR" -xzf "$DOWNLOAD_FILE"
fi

# Clean up
rm "$DOWNLOAD_FILE"

echo -e "${GREEN}✓ Installation complete${NC}\n"

# Verify installation
NEW_VERSION=$(go version 2>/dev/null | awk '{print $3}' | sed 's/go//' || echo "error")

if [ "$NEW_VERSION" = "$LATEST_VERSION" ]; then
    echo -e "${GREEN}✓ Successfully updated Go to version ${LATEST_VERSION}${NC}"
else
    echo -e "${YELLOW}⚠ Go was installed but version mismatch detected${NC}"
    echo -e "Expected: ${LATEST_VERSION}, Got: ${NEW_VERSION}"
    echo -e "\nYou may need to update your PATH:"
    echo -e "export PATH=\$PATH:${INSTALL_PATH}/bin"
fi

# Check PATH configuration
echo -e "\n${BLUE}Checking PATH configuration...${NC}"
if ! echo "$PATH" | grep -q "${INSTALL_PATH}/bin"; then
    echo -e "${YELLOW}⚠ ${INSTALL_PATH}/bin is not in your PATH${NC}"
    echo -e "\nAdd this to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    echo -e "${GREEN}export PATH=\$PATH:${INSTALL_PATH}/bin${NC}"
fi

echo -e "\n${GREEN}Done!${NC}"
go version
