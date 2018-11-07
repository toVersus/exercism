#!/usr/bin/env bash

set -o errexit

INPUT=$1

reverse_string() {
    rev=""
    len=${#INPUT}
    for (( i=$len-1 ; i>=0; i-- )) ; do
        rev+="${INPUT:$i:1}";
    done

    echo $rev
}

reverse_string