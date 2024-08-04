#!/bin/sh

if [ "$1" = 'configure' ] && [ -z "$2" ]; then
    printf "\033[32m[switchboard][info] Installing switchboard agent...\033[0m\n"
    systemctl daemon-reload
    systemctl unmask switchboard-agent
    systemctl preset switchboard-agent
    systemctl enable switchboard-agent
    systemctl restart switchboard-agent
fi
