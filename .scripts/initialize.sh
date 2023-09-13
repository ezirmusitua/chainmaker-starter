#!/bin/sh
VERSION=v2.3.0

git submodule add https://git.chainmaker.org.cn/chainmaker/management-backend.git management-backend
pushd management-backend
  git checkout tags/v2.3.0 -b v2.3.0
popd
