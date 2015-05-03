# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  # Setup virtual machine box. This VM configuration code is always executed.
  config.vm.box     = "ubuntu"
  config.vm.box_url = "http://files.vagrantup.com/precise64.box"

  # use ssh agent forwarding so we effectively use our host ssh settings
  config.ssh.forward_agent = true

  # modify the memory we allocate
  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", "2056"]
  end

  config.vm.network "private_network", ip: "192.168.50.4"

  # web port we use to access things
  config.vm.network "forwarded_port", guest: 8080, host: 8080, auto_correct: false # web
  # config.vm.network "forwarded_port", guest: 8181, host: 8181, auto_correct: false # browser-sync ui - doesn't work, attempts to route to internal port. changable?
  config.vm.network "forwarded_port", guest: 9090, host: 9090, auto_correct: false # go-convey test server

  # run this script to set things up. 
  config.vm.provision "shell", path: "provision/setup.sh"

  # mount this project's 'tmp' folder into the vagrant instance as osx
  config.vm.synced_folder ".", "/home/vagrant/clikr", :type => "rsync"
end
