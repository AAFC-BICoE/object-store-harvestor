#!/bin/bash

# exit when any command fails
set -e

###############################################################
# Validation
# This is a guard to prevent running incorrect golang versions
###############################################################
required_golang="1.16"
v=`go version | { read _ _ v _; echo ${v#go}; }`
if [[ $v =~ $required_golang ]]; then
   echo "go version is good and is : $v"
else
   echo "= = = = = = = = = = = = = = = = = = = = = = = "
   echo "This application requires golang version 1.16+"
   echo "Current golang version is $v"
   echo "Please upgrade golang to version 1.16+"
   echo "= = = = = = = = = = = = = = = = = = = = = = = "
   exit 1
fi
###############################################################

# clean up
rm -fr /tmp/db-test
echo "/tmp/db-test has been removed ..."
rm -fr /tmp/data-test
echo "/tmp/data-test has been removed ..."
rm -fr /tmp/log-test
echo "/tmp/log-test has been removed ..."
# prepare test data set
mkdir -p /tmp/db-test/
echo "/tmp/db-test has been set ..."
cp -R ../data-test /tmp/
echo "/tmp/data-test has been set ..."
mkdir -p /tmp/log-test/
echo "/tmp/log-test has been set ..."

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
# clean up after test
# rm -fr /tmp/db-test/harvestor.db
echo "DB /tmp/db-test/harvestor.db has been removed ..."
# building the artifact
echo "||| building the artifact ..."
go build -o main .
# running ...
./main ./harvestor_config.yml
####################
