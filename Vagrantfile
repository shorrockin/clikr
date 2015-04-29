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

  # web port we use to access things
  config.vm.network "forwarded_port", guest: 8080, host: 8080, auto_correct: false

  # run this script to set things up. 
  config.vm.provision "shell", path: "provision/setup.sh"
end
