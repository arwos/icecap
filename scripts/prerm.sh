#!/bin/bash


if [ -f "/etc/systemd/system/icecap.service" ]; then
    systemctl stop icecap
    systemctl disable icecap
    systemctl daemon-reload
fi
