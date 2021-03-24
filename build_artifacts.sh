#!/bin/bash

# exit when any command fails
set -e

# this is a local image
docker_repo_target="aafc-bicoe/object-store-harvestor-local"
# building docker image
docker build -t $docker_repo_target -f deployment-local/Dockerfile-Artifacts .
# bring up/update containers
docker-compose -f deployment-local/docker-compose.yml up
# clean up
rm -fr artifacts/*
# building linux & windows targets
docker cp aafc-bicoe-object-store-harvestor-local:/app/main artifacts/harvestor-linux-amd64
#docker cp aafc-bicoe-object-store-harvestor-local:/app/main.exe artifacts/harvestor-windows-amd64.exe
