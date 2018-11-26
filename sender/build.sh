#!/usr/bin/env bash
set -e

# build go application
go build

# build docker image and push
docker build -t anfra/sender:latest .
docker push anfra/sender

