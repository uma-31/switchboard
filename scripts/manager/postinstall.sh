#!/bin/sh

if [ "$1" = 'configure' ] && [ -z "$2" ]; then
    printf "\033[32m[switchboard][info] Installing switchboard manager...\033[0m\n"
    systemctl daemon-reload
    systemctl unmask switchboard-manager
    systemctl preset switchboard-manager
    systemctl enable switchboard-manager
    systemctl restart switchboard-manager
fi
