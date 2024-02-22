#!/bin/bash


if [ -f "/etc/systemd/system/icecap.service" ]; then
    systemctl start icecap
    systemctl enable icecap
    systemctl daemon-reload
fi
