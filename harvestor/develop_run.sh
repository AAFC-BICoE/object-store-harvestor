#!/bin/bash

# exit when any command fails
set -e

# clean up
rm -fr /tmp/db-test
echo "/tmp/db-test has been removed ..."
rm -fr /tmp/data-test
echo "/tmp/data-test has been removed ..."
# prepare test data set
mkdir -p /tmp/db-test/
echo "/tmp/db-test has been set ..."
cp -R ../data-test /tmp/
echo "/tmp/data-test has been set ..."

####################
### D E V  R U N ###
# going modules
echo "||| pulling modules ..."
go mod tidy
# init build 
echo "||| init build ..."
go build -o main .
# provide coverprofile
echo "||| setup coverprofile ..."
go test ./... -v -coverprofile /tmp/cover.out
# generate coverage html file
echo "||| testing ..."
go tool cover -html=/tmp/cover.out -o /tmp/cover.html
# building the artifact
echo "||| building the artifact ..."
go build -o main .
# running ...
./main ./harvestor_config.yml
####################
