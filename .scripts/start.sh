#!/bin/sh
cd /src/chainmaker
nerdctl compose down && nerdctl compose up -d --build

ln -sf /src/chainmaker/chainmaker.conf /etc/nginx/sites-enabled/chainmaker.conf

nginx -t && nginx -s reload

