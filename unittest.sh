#!/bin/bash

export PGMGR_CONFIG_FILE=".pgmgr-test.json"

pgmgr db drop

pgmgr db create

pgmgr db migrate

revel test github.com/elirenato/golangseed test