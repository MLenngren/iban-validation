#!/bin/bash
sudo apt-get update -y &&
sudo apt-get install -y \
apt-transport-https \
ca-certificates \
curl \
gnupg-agent \
software-properties-common \
docker.io &&
sudo apt-get update -y &&
sudo usermod -aG docker ubuntu