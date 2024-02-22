#!/bin/bash


if ! [ -d /var/lib/icecap/ ]; then
    mkdir /var/lib/icecap
fi

if [ -f "/etc/systemd/system/icecap.service" ]; then
    systemctl stop icecap
    systemctl disable icecap
    systemctl daemon-reload
fi
