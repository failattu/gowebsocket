# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.define "go1" do |go1|
    go1.vm.box = "ubuntu/trusty64"
    go1.vm.network "public_network"
    go1.vm.provider "virtualbox" do |v|
  	  v.memory = 2048
  	  v.cpus = 2
    end
	go1.vm.provision "shell" do |s|
     s.path = "deploy.sh"
     s.privileged = true
    end
  end
end
