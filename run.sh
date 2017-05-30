#!/bin/bash
profile="dev"
if [ "x$1" != "x" ]; then
    profile=$1
fi
goVersion="$(go version)"
echo "#######################"
echo "Go version: $goVersion"
echo "Using profile: $profile"
echo "#######################"
revel run github.com/elirenato/golangseed $profile