#!/bin/bash
 
echo "Provisioning virtual machine..."

echo "Installing some basic tools"
sudo apt-get -y install curl 

echo "Installing Docker"
wget -qO- https://get.docker.com/ | sh

echo "Installing Docker Compose"
curl -L https://github.com/docker/compose/releases/download/1.2.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
sudo usermod -a -G docker vagrant

echo "Cloning the clikr repo"
ssh-keyscan github.com >> ~/.ssh/known_hosts
git clone git@github.com:shorrockin/clikr.git
sudo chown vagrant:vagrant -R clikr
