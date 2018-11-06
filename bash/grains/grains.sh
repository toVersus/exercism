#!/usr/bin/env bash

set -o errexit
set -o nounset

INPUT=$1

grains() {
  if [ $INPUT = "total" ]; then
  	echo "2^64-1" | bc
  elif (( $INPUT < 1 || $INPUT > 64 )) ; then
    echo "Error: invalid input"
    exit 1
  else
    Count=$(( 2**($INPUT-1) ))
    echo ${Count#-}
  fi
}

grains
