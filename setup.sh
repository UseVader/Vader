#!/bin/bash

INSTALL_DIR="/usr/local/bin"
EXECUTABLE_NAME="vader"  # Replace with your CLI name
EXECUTABLE_PATH="$INSTALL_DIR/$EXECUTABLE_NAME"
USE_SUDO="false"
OS=""
ARCH=""

# Colors for output
RED='\033[0;31m'
PURPLE='\033[0;35m'
GREEN='\033[0;32m'
NC='\033[0m'

setSystem() {
    ARCH=$(uname -m)
    case $ARCH in
        i386|i686) ARCH="386" ;;    # Changed for Go naming convention
        x86_64) ARCH="amd64" ;;     # Changed for Go naming convention
        armv6*) ARCH="arm" ;;       # Changed for Go naming convention
        armv7*) ARCH="arm" ;;       # Changed for Go naming convention
        aarch64*) ARCH="arm64" ;;
    esac

    OS=$(echo uname|tr '[:upper:]' '[:lower:]')
    if [ "$OS" = "linux" ]; then
        USE_SUDO="true"
    fi
    if [ "$OS" = "darwin" ] && [[ "$(uname -a)" = ARM64 ]]; then
        USE_SUDO="true"
    fi
}

runAsRoot() {
    local CMD="$*"
    if [ "$USE_SUDO" = "true" ]; then
        printf "${PURPLE}We need sudo access to install to $INSTALL_DIR ${NC}\n"
        CMD="sudo $CMD"
    fi
    $CMD
}

downloadBinary() {
    BINARY_NAME="${EXECUTABLE_NAME}${OS}${ARCH}"
    # Replace with your GitHub repository URL
    GITHUB_URL="https://github.com/username/repo/releases/latest/download/$BINARY_NAME"
    printf "${PURPLE}Downloading $BINARY_NAME...${NC}\n"
    curl -L --progress-bar --output "$EXECUTABLE_NAME-tmp" "$GITHUB_URL"
    if [ $? -ne 0 ]; then
        printf "${RED}Failed to download binary... ${NC}\n"
        exit 1
    fi
}

install() {
    chmod +x "$EXECUTABLE_NAME-tmp"
    if [ $? -ne 0 ]; then
        printf "${RED}Failed to set permissions... ${NC}\n"
        exit 1
    fi

    runAsRoot mv "$EXECUTABLE_NAME-tmp" "$EXECUTABLE_PATH"
    if [ $? -ne 0 ]; then
        printf "${RED}Failed to move binary to $INSTALL_DIR... ${NC}\n"
        exit 1
    fi
}

printf "${PURPLE}Installing ${EXECUTABLE_NAME}...${NC}\n\n"

setSystem
downloadBinary
install
printf "\n${GREEN}Successfully installed to:${NC} $EXECUTABLE_PATH\n"