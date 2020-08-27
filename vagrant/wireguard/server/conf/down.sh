#!/bin/bash

# Routing
iptables -D INPUT -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
iptables -D FORWARD -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
iptables -D INPUT -p udp -m udp --dport 51820 -m conntrack --ctstate NEW -j ACCEPT
iptables -D INPUT -s 10.200.200.0/24 -p tcp -m tcp --dport 53 -m conntrack --ctstate NEW -j ACCEPT
iptables -D INPUT -s 10.200.200.0/24 -p udp -m udp --dport 53 -m conntrack --ctstate NEW -j ACCEPT
iptables -D FORWARD -i wg0 -o wg0 -m conntrack --ctstate NEW -j ACCEPT

# NAT
iptables -t nat -D POSTROUTING -s 10.200.200.0/24 -o eth0 -j MASQUERADE
