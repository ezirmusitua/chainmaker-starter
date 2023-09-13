#!/bin/sh
REMOTE=$1

ssh $REMOTE -t "mkdir -p /src/chainmaker && mkdir -p /data/chainmaker/{mgmt,explorer}"

rsync -av . $REMOTE:/src/chainmaker \
  --exclude=.DS_Store \
  --exclude=.git \
  --exclude=build \
  --exclude=management-web \
  --exclude=chainmaker-explorer-web \
  --exclude=node_modules

ssh $REMOTE -t < .scripts/start.sh
