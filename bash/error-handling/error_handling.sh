#!/usr/bin/env bash

error_no_arg() {
    echo "Usage: ./error_handling <greetee>"
    exit 1
}

hello() {
    if [ $# -eq 0 ]; then
        error_no_arg
    elif [ $# -eq 1 ]; then
        echo "Hello, $1"
        exit 0
    fi
    exit 1
}

hello "$@"
