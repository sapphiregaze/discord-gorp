#!/bin/bash

# Variables
BINARY_NAME="discord-gorp"
BINARY_PATH="$HOME/.local/bin/$BINARY_NAME"
SERVICE_NAME="discord-gorp.service"
SERVICE_PATH="$HOME/.config/systemd/user/$SERVICE_NAME"

# Stop and disable the systemd service if it is running
echo "Stopping and disabling the $SERVICE_NAME service..."
systemctl --user stop "$SERVICE_NAME" 2>/dev/null
systemctl --user disable "$SERVICE_NAME" 2>/dev/null

# Remove the systemd service file
if [ -f "$SERVICE_PATH" ]; then
  echo "Removing the systemd service file at $SERVICE_PATH..."
  rm "$SERVICE_PATH"
else
  echo "No systemd service file found at $SERVICE_PATH."
fi

# Reload systemd daemon to remove the service reference
systemctl --user daemon-reload

# Remove the binary from ~/.local/bin
if [ -f "$BINARY_PATH" ]; then
  echo "Removing the binary at $BINARY_PATH..."
  rm "$BINARY_PATH"
else
  echo "No binary found at $BINARY_PATH."
fi

echo "Uninstallation complete."
