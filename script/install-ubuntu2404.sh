#!/bin/bash
# Install the latest versions of Docker Engine and Docker Compose on Ubuntu 24.04
# Exit immediately if a command exits with a non-zero status
set -e

echo "1/6: Updating package lists..."
sudo apt-get update

echo "2/6: Installing prerequisites..."
sudo apt-get install -y ca-certificates curl

echo "3/6: Adding Docker's official GPG key..."
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

echo "4/6: Adding Docker repository to Apt sources..."
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

echo "5/6: Installing Docker Engine and Docker Compose..."
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

echo "6/6: Adding user to the docker group..."
# Uses SUDO_USER if run with sudo, otherwise falls back to USER
TARGET_USER="${SUDO_USER:-$USER}"
sudo usermod -aG docker "$TARGET_USER"

echo "========================================"
echo "Installation complete!"
echo "========================================"
docker --version
docker compose version
echo ""
echo "IMPORTANT: To apply the docker group changes, you must log out and log back in,"
echo "or run this command right now: newgrp docker"
