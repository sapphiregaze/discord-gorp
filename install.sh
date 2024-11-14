#!/bin/bash

# Variables
BINARY_NAME="discord-gorp"
BINARY_PATH="$HOME/.local/bin/$BINARY_NAME"
SERVICE_NAME="discord-gorp.service"
SERVICE_PATH="$HOME/.config/systemd/user/$SERVICE_NAME"

# Check if `make` is installed
if ! command -v make &> /dev/null; then
  echo "Error: 'make' is not installed. Please install it and try again."
  exit 1
fi

# Check if `go` is installed
if ! command -v go &> /dev/null; then
  echo "Error: 'go' is not installed. Please install it and try again."
  exit 1
fi

# Build the binary
echo "Building binary with make..."
if ! make; then
  echo "Error: Failed to build the binary."
  exit 1
fi

# Ensure ~/.local/bin exists and is in PATH
mkdir -p "$HOME/.local/bin"
export PATH="$HOME/.local/bin:$PATH"

# Check if the binary file exists in the current directory
if [ ! -f "./$BINARY_NAME" ]; then
  echo "Error: $BINARY_NAME binary not found in the current directory."
  exit 1
fi

# Copy the binary to ~/.local/bin
cp "./$BINARY_NAME" "$BINARY_PATH"
chmod +x "$BINARY_PATH"
echo "Copied $BINARY_NAME to $BINARY_PATH"

# Ensure systemd user directory exists
mkdir -p "$(dirname "$SERVICE_PATH")"

# Create the systemd service file
cat <<EOF > "$SERVICE_PATH"
[Unit]
Description=Discord Gorp Service
After=network.target

[Service]
ExecStart=$BINARY_PATH
Restart=on-failure

[Install]
WantedBy=default.target
EOF

echo "Created systemd service file at $SERVICE_PATH"

# Reload systemd user daemon to register the new service
systemctl --user daemon-reload

# Enable and start the service
systemctl --user enable --now "$SERVICE_NAME"

echo "Service $SERVICE_NAME has been installed and started."
echo "You can check the status with: systemctl --user status $SERVICE_NAME"
