#!/bin/bash

serviceName='Wi-Fi'
proxyDomain='117.50.96.61'
proxyPort=8080

if [ "$1" == "off" ] ; then
    echo "handshakeproxy off"
    networksetup -setwebproxystate 'Wi-Fi' off
    networksetup -setsecurewebproxystate 'Wi-Fi' off
else    
    echo "handshakeproxy on"
    networksetup -setsecurewebproxy $serviceName $proxyDomain $proxyPort 
    networksetup -setwebproxy $serviceName $proxyDomain $proxyPort 
fi


