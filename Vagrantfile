# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

$jekyll_script = <<SCRIPT
apt-get -y install curl
curl -L https://get.rvm.io | bash -s stable --ruby=2.0.0
gem install bundler
cd /vagrant
bundle install
SCRIPT

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "hashicorp/precise64"
  config.vm.network "forwarded_port", guest: 4000, host: 4000

  config.vm.provision "shell", inline: $jekyll_script
end
