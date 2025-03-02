#!/bin/bash

OS=$(uname -s)
ARCH=$(uname -m)

if [ "$OS" = "Linux" ] && [ "$ARCH" = "x86_64" ]; then
  BINARY_URL="https://github.com/chcknlegwill/chcknlegwill-cli/releases/download/v1.0.2/chcknlegwill-cli-linux-x86_64-v1.0.2" #linux version
elif [ "$OS" = "Darwin" ] && [ "$ARCH" = "arm64" ]; then
  BINARY_URL="https://github.com/chcknlegwill/chcknlegwill-cli/releases/download/v1.0.2/chcknlegwill-cli-macos-amd64-v1.0.2" #macos version
else
  echo "Unsupported OS or architecture: $OS $ARCH"
  exit 1
fi

echo "Downloading chcknlegwill-cli binary..." #make a loading animation like with npm
curl -L "$BINARY_URL" -o chcknlegwill-cli

echo "Making the executable..."
chmod +x chcknlegwill-cli


# Check if binary already exists in /usr/local/bin
if [ -f "/usr/local/bin/chcknlegwill-cli" ]; then
  echo "Warning: chcknlegwill-cli already exists in /usr/local/bin. Overwrite? (y/n)"
  read -r response
  if [ "$response" != "y" ] && [ "$response" != "Y" ]; then
    echo "Installation aborted."
    exit 1
  fi
  sudo mv /usr/local/bin/chcknlegwill-cli /usr/local/bin/chcknlegwill-cli.old
fi

# Install to /usr/local/bin with sudo
echo "Installing to /usr/local/bin..."
sudo mv chcknlegwill-cli /usr/local/bin/

if ! echo "$PATH" | grep -q "/usr/local/bin"; then
  echo "Note: Ensure /usr/local/bin is in your PATH. You may need to log out / in or run export PATH\"usr/local/bin:$PATH\""
fi

echo "Installation complete!"
echo "To use chcknlegwill cli, simply type it into the CLI."
