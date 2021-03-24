#!/bin/bash

# exit when any command fails
set -e

# out folders
dist_base="aafc-bicoe"
dist="harvestor"
# this is a local image
docker_repo_target="aafc-bicoe/object-store-harvestor-local"
# building docker image
docker build -t $docker_repo_target -f deployment-local/Dockerfile-Artifacts .
# bring up/update containers
docker-compose -f deployment-local/docker-compose.yml up
# clean up
rm -fr $dist_base/*
# building linux & windows targets
mkdir -p $dist_base/$dist
docker cp aafc-bicoe-object-store-harvestor-local:/app/main $dist_base/$dist/harvestor-linux-amd64
cp harvestor/harvestor_config.yml $dist_base/$dist/
tar cvf $dist_base/aafc-bicoe-harvestor-linux-amd64.v0.1.tar $dist_base/$dist
#cleanup
rm -fr $dist_base/$dist

#docker cp aafc-bicoe-object-store-harvestor-local:/app/main.exe artifacts/harvestor-windows-amd64.exe
