#!/bin/bash
# Install wireguard
sudo echo "deb http://deb.debian.org/debian/ unstable main" > /etc/apt/sources.list.d/unstable.list
sudo printf 'Package: *\nPin: release a=unstable\nPin-Priority: 90\n' > /etc/apt/preferences.d/limit-unstable
sudo apt update -y && sudo apt install -y wireguard

# Install config
sudo cp /home/vagrant/wg0.conf /etc/wireguard/