#!/bin/bash

# Install golang
echo "Setting up workspace"
sudo chown vagrant workspace
mkdir workspace/sdks workspace/sdks/go_1.14.7
cd workspace/sdks/go_1.14.7

# Download and unpack
echo "Downloading and unpacking go 1.14.7"
wget https://golang.org/dl/go1.14.7.linux-amd64.tar.gz -o /dev/null
tar -xf go1.14.7.linux-amd64.tar.gz

sudo ln -s /home/vagrant/workspace/sdks/go_1.14.7/go /usr/local/go

# Set env properties and path
echo "Setting up the PATH"
echo '# Custom settings' >> /home/vagrant/.profile
echo 'export GOROOT=/usr/local/go' >> /home/vagrant/.profile
echo 'export GOPATH=$HOME/go' >> /home/vagrant/.profile
echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> /home/vagrant/.profile
source /home/vagrant/.profile