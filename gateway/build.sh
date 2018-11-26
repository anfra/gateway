#!/usr/bin/env bash
set -e

# build go application
go build

# build docker image and push
docker build -t anfra/gateway:latest .
docker push anfra/gateway

