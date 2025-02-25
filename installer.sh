#!/bin/bash

OS=$(uname -s)
ARCH=$(uname -m)

if [ "SOS" = "Linux" ] && ["$ARCH" = "x86_64"]; then
  BINARY_URL="https://github.com/chcknlegwill/chcknlegwill-cli/releases/download/First/chcknlegwill-cli-linux-x86_64"  #linux version #finish this
elif [ "SOS" = "Darwin" ] && [ "$ARCH" = "arm64" ]; then
  BINARY_URL="https://github.com/chcknlegwill/chcknlegwill-cli/releases/download/First/chcknlegwill-cli-macos-x86_64" #macos version
else
  echo "Unsupported OS or architecture: $OS $ARCH"
  exit 1
fi

echo "Downloading chcknlegwill-cli binary..." #make a loading animation like with npm
curl -L "$BINARY_URL" -o chcknlegwill-cli

echo "Making the executable..."
chmod +x chcknlegwill-cli

if [ ! -d "$HOME/bin" ]; then
  echo "Creating ~/bin directory..."
  mkdir "$HOME/bin"
fi

echo "Installation complete!"
echo "To use chcknlegwill cli, simply type it into the CLI."
