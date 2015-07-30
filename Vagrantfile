# -*- mode: ruby -*-
# vi: set ft=ruby :


disable_firewall = "service iptables stop"

Vagrant.configure("2") do |config|
  config.vm.define 'kudzu' do |box_config|
    box_config.vm.box = 'puppetlabs/centos-6.6-64-nocm'
    box_config.vm.hostname = 'kudzu.example.com'

    box_config.vm.provider :vmware_fusion do |v|
      v.vmx["memsize"]  = '512'
      v.vmx["numvcpus"] = '1'
    end
  end

  config.vm.provision "shell", inline: disable_firewall
end
