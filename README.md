# object-store-harvestor

* TODO setup git repo on https://travis-ci.org to cover a build status (it's free) and may be more repo badges

One Paragraph of project description goes here


## Contents

* TODO High level design (may be goes to a design folder)

- [Object Store Harvestor](#object-store-harvestor)
  - [Docs](#docs)
  - [Contents](#contents)
  - [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Local development](#local-development)
  - [Running the tests](#running-the-tests)

## Docs
  - [Design](doc/design.md)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- Install [Docker](https://docs.docker.com/get-docker/)
- Install [docker-compose](https://docs.docker.com/compose/)

### Local development

- Clone the repo: object-store-harvestor
- cd to repo object-store-harvestor
- run the following command to start a docker container :
```
deployment-local/run_docker_containers.sh
```
- check the logs of the aafc-bicoe-object-store-harvestor-local conatner for results
```
docker logs aafc-bicoe-object-store-harvestor-local
```
## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* TODO

## Versioning

* TODO

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
