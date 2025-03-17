#!/bin/bash


show_loading() {
  local pid=$1
  local delay=0.1
  local spin='-\|/'

  while kill -0 "$pid" 2>/dev/null; do
    for i in $(seq 0 3); do
    echo -ne "\r${spin:$i:1} Downloading chcknlegwill-cli binary..."
    sleep $delay
    done
  done
  echo -e "\rDone downloading chcknlegwill-cli binary.\n"
}

OS=$(uname -s)
ARCH=$(uname -m)

if [ "$OS" = "Linux" ] && [ "$ARCH" = "x86_64" ]; then
  BINARY_URL="https://github.com/chcknlegwill/chcknlegwill-cli/releases/download/v1.0.4/chcknlegwill-cli-linux-x86_64-v1.0.4" #linux version
elif [ "$OS" = "Darwin" ] && [ "$ARCH" = "arm64" ]; then
  BINARY_URL="https://github.com/chcknlegwill/chcknlegwill-cli/releases/download/v1.0.4/chcknlegwill-cli-macos-amd64-v1.0.4" #macos version
else
  echo "Unsupported OS or architecture: $OS $ARCH"
  exit 1
fi

echo "Starting download..."
curl -L "$BINARY_URL" -o chcknlegwill-cli &
curl_pid=$!
show_loading $curl_pid
wait $curl_pid
if [$? -ne 0 ]; then
  echo "Error: Failed to download the binary."
  rm -f chcknlegwill-cli
  exit 1
fi

echo "Making executable..."
if ! chmod +x chcknlegwill-cli; then
  echo "Error: Failed to make the executable."
  rm -f chcknlegwill-cli
  exit 1
fi

#/usr/local/bin is protected so you will need sudo, but makes it so you can call the 
# tool anywhere e.g. in ~ (/home/$USER/) or in ~/Documents/project/ for ease of use
INSTALL_DIR="/usr/local/bin"
TARGET_PATH="$INSTALL_DIR/chcknlegwill-cli"

# Check if binary already exists in /usr/local/bin
if [ -f "$TARGET_PATH" ]; then
    while true; do
        echo "Warning: chcknlegwill-cli already exists in $INSTALL_DIR. Overwrite? (y/n/r)"
        echo "  (y) Yes, overwrite the existing file"
        echo "  (n) No, abort installation"
        echo "  (r) Rename the existing file and continue"
        read -r response
        case "$response" in
            [Yy]*)
                break
                ;;
            [Nn]*)
                echo "Installation aborted."
                rm -f chcknlegwill-cli
                exit 1
                ;;
            [Rr]*)
                if ! sudo mv "$TARGET_PATH" "$TARGET_PATH.old"; then
                    echo "Error: Failed to rename existing file. Installation aborted."
                    rm -f chcknlegwill-cli
                    exit 1
                fi
                break
                ;;
            *)
                echo "Invalid input. Please enter 'y', 'n', or 'r'."
                ;;
        esac
    done
fi

# Install to /usr/local/bin with sudo (required)
echo "Installing to $INSTALL_DIR..."
if ! sudo mv chcknlegwill-cli "$TARGET_PATH"; then
    echo "Error: Failed to install chcknlegwill-cli to $INSTALL_DIR with sudo. Installation aborted."
    rm -f chcknlegwill-cli
    exit 1
fi

# Check if /usr/local/bin is in PATH and provide instructions if not
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
    echo "Note: Ensure $INSTALL_DIR is in your PATH. You may need to log out/in or run:"
    echo "      export PATH=\"/usr/local/bin:$PATH\""
fi

echo "Installation complete!"
echo "To use chcknlegwill-cli, simply type 'chcknlegwill-cli' in the terminal."