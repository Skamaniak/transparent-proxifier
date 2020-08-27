#!/bin/bash

# Setup forwarding
echo "net.ipv4.ip_forward=1" >>/etc/sysctl.conf
sysctl -p
echo 1 >/proc/sys/net/ipv4/ip_forward # maybe unnecessary

wg-quick up wg0
