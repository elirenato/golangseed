#!/bin/bash

goVersion="$(go version)"
echo "#######################"
echo "Go version: $goVersion"
echo "#######################"

export PGMGR_CONFIG_FILE=".pgmgr-test.json"

pgmgr db drop

pgmgr db create

pgmgr db migrate


if [ "x$1" != "x" ]; then
    revel test github.com/elirenato/golangseed test $1
else 
    revel test github.com/elirenato/golangseed test
fi

