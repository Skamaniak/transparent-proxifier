# Install on Ubuntu 20.04

TProxy on linux - https://www.kernel.org/doc/html/latest/networking/tproxy.html (https://www.kernel.org/doc/Documentation/networking/tproxy.txt)



# Draft notes
Need to be installed on host machine 
`vagrant plugin install vagrant-vbguest`

#### Setup nftables (https://www.kernel.org/doc/html/latest/networking/tproxy.html)
```
sudo iptables -t mangle -N DIVERT
sudo iptables -t mangle -A PREROUTING -p tcp -m socket -j DIVERT
sudo iptables -t mangle -A DIVERT -j MARK --set-mark 1
sudo iptables -t mangle -A DIVERT -j ACCEPT

sudo ip rule add fwmark 1 lookup 100
sudo ip route add local 0.0.0.0/0 dev lo table 100

sudo iptables -t mangle -A PREROUTING -p tcp --dport 80 -j TPROXY --tproxy-mark 0x1/0x1 --on-port 8080
```


Santa's little helpers
```
sudo iptables -t mangle --list
sudo ip rule
ip route show table all
ip route list table local
```

## Sources
#### TProxy
 * https://powerdns.org/tproxydoc/tproxy.md.html
 * https://www.kernel.org/doc/html/latest/networking/tproxy.html
#### Wireguard and setup
 * https://softwaretester.info/simple-vpn-via-wireguard/
 * https://www.ckn.io/blog/2017/11/14/wireguard-vpn-typical-setup/
