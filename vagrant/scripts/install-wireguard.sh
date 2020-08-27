#!/bin/bash
# Install wireguard
echo "deb http://deb.debian.org/debian/ unstable main" >/etc/apt/sources.list.d/unstable.list
printf 'Package: *\nPin: release a=unstable\nPin-Priority: 90\n' >/etc/apt/preferences.d/limit-unstable
apt update -y && sudo apt install -y wireguard

# Install config
cp -r /tmp/wg-conf/* /etc/wireguard/
# Make scripts (if any) executable
find /etc/wireguard/ -type f -iname "*.sh" -exec chmod +x {} \;
# Replace the windows line endings (Duh!) with the linux ones if there are any
find /etc/wireguard/ -type f -iname "*.sh" -exec sed -i 's/\r$//' {} \;
