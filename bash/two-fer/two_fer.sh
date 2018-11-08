#!/usr/bin/env bash

set -o errexit

INPUT=$1

two_fer() {
        echo "One for ${INPUT:-you}, one for me."
}

two_fer
