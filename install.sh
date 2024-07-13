#!/bin/bash

# If no version is specified as a command line argument, fetch the latest version.
if [ -z "$1" ]; then
    VERSION=$(curl -s https://api.github.com/repos/tttol/mergectl/releases/latest | grep -o '"tag_name": *"[^"]*"' | sed 's/"tag_name": *"//' | sed 's/"//')
    if [ -z "$VERSION" ]; then
        echo "Failed to fetch the latest version"
        exit 1
    fi
else
    VERSION=$1
fi

OS=$(uname -s)
ARCH=$(uname -m)
URL="https://github.com/tttol/mergectl/releases/download/${VERSION}/mergectl_${OS}_${ARCH}.tar.gz"

echo "Start to install."
echo "VERSION=$VERSION, OS=$OS, ARCH=$ARCH"
echo "URL=$URL"

TMP_DIR=$(mktemp -d)
curl -L $URL -o $TMP_DIR/mergectl.tar.gz
tar -xzvf $TMP_DIR/mergectl.tar.gz -C $TMP_DIR
sudo mv $TMP_DIR/mergectl /usr/local/bin/mergectl
sudo chmod +x /usr/local/bin/mergectl

rm -rf $TMP_DIR

if [ -f "/usr/local/bin/mergectl" ]; then
  echo "[SUCCESS] mergectl $VERSION has been installed to /usr/local/bin"
else
  echo "[FAIL] mergectl $VERSION is not installed."
fi