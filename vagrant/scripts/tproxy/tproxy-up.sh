#!/bin/bash

# https://powerdns.org/tproxydoc/tproxy.md.html
iptables -t mangle -N DIVERT
iptables -t mangle -A PREROUTING -p tcp -m socket -j DIVERT
iptables -t mangle -A DIVERT -j MARK --set-mark 1
iptables -t mangle -A DIVERT -j ACCEPT

# Setup policy
ip rule add fwmark 1 lookup 100
ip route add local 0.0.0.0/0 dev lo table 100

# TProxy anything going to destination :80 and :443
iptables -t mangle -A PREROUTING -p tcp --dport 80 -j TPROXY --tproxy-mark 0x1/0x1 --on-port 1080
iptables -t mangle -A PREROUTING -p tcp --dport 443 -j TPROXY --tproxy-mark 0x1/0x1 --on-port 1081
