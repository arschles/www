#!/bin/bash

set -e

sudo apt install -y jq

echo "all file additions:"
ADDED_FILES=$(cat $HOME/files_added.json)
echo
