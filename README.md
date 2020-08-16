## transparent-proxy
The app is able to pick up tcp connection that are destined to somewhere else and redirect them to a HTTP proxy.
The intended usage is that the Linux transparent proxy is configured to pickup anything destined to :80 and :443 and call this app with it
the app then
 * Bridges the :80 tcp to an HTTP proxy
 * Forgex CONNECT request for :443 tcp, sends it to the HTTP proxy and then bridges
 
Effectively this app listens to any HTTP/HTTPS traffic, slurps it and forces it to go through an HTTP proxy even though
the originator of those requests doesn't have any proxy configured

## How to use?
Project contains a Vagrant file that spins up 2 machines (client, server). These machines are connected via Wireguard VPN.

1) Start the Vagrant env
2) SSH to the "server" machine
3) go to workspace/repositories/transparent-proxifier
4) run `go build`
5) run sudo `./transparent-proxifier`
6) in another terminal SSH to the "server" machine as well
7) go to workspace/repositories/ssl-decrypter (forward proxy + SSL decryption)
8) run `go build`
9) run `./ssl-decrypter`
10) in another terminal SSH to the "client"
11) run `curl google.com` (you should see request going through)
12) run `curl https://google.com` -- this will return SSL error (due to SSL decryption)
13) [install the self-signed CA into the system](https://askubuntu.com/questions/73287/how-do-i-install-a-root-certificate) ("client" box) - you can find it in "workspace/repositories/ssl-decryption/assets" on the "server" machine
14) repeat step 12) and you should get response without the SSL error

## Troubleshooting
* `sudo wg` on any of the machine to troubleshoot the VPN connection

## Sources
#### TProxy
 * https://powerdns.org/tproxydoc/tproxy.md.html
 * https://www.kernel.org/doc/html/latest/networking/tproxy.html
#### Wireguard and setup
 * https://softwaretester.info/simple-vpn-via-wireguard/
 * https://www.ckn.io/blog/2017/11/14/wireguard-vpn-typical-setup/
