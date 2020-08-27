#!/bin/bash

# https://powerdns.org/tproxydoc/tproxy.md.html
iptables -t mangle -D PREROUTING -p tcp -m socket -j DIVERT
iptables -t mangle -D DIVERT -j MARK --set-mark 1
iptables -t mangle -D DIVERT -j ACCEPT
iptables -t mangle -X DIVERT


# Remove policy
ip rule del fwmark 1 lookup 100
ip route del local 0.0.0.0/0 dev lo table 100

# TProxy anything going to destination :80 and :443
iptables -t mangle -D PREROUTING -p tcp --dport 80 -j TPROXY --tproxy-mark 0x1/0x1 --on-port 1080
iptables -t mangle -D PREROUTING -p tcp --dport 443 -j TPROXY --tproxy-mark 0x1/0x1 --on-port 1081
