⛔️ This project is now archived - No Maintenance Intended

# object-store-harvestor

## Contents

- [Object Store Harvestor](#object-store-harvestor)
  - [Docs](#docs)
  - [Prerequisites](#prerequisites)
  - [Local development](#local-development)
  - [Running the tests](#running-the-tests)
  - [Built With](#built-with)
  - [Artifacts](#artifacts)
  - [Versioning](#versioning)
  - [Deployment](#deployment)

## Docs
  - [Application Design](doc/application.md)
  - [Application Diagram](doc/design-diagram-v0.02.pdf)
  - [Persistent Storage Design](doc/persistent-storage.md)
  - [Persistent Storage Diagram](doc/sqlite-db-diagram-v0.01.pdf)
  - [Application Config](doc/application-config.md)

## Local development
 - `prerequisites`
    - Install [Docker](https://docs.docker.com/get-docker/)
    - Install [docker-compose](https://docs.docker.com/compose/)
    - Setup [dina-local-deployment](https://github.com/AAFC-BICoE/dina-local-deployment)

    ##### Docer local development
    - clone the repo: object-store-harvestor
    - cd to repo object-store-harvestor
    - run the following command to start a docker container :
        ```
        ./run_docker_containers.sh
        ```
    ##### local development (linux)
    - clone the repo: object-store-harvestor
    - cd to repo object-store-harvestor
    - cd harvestor
    - run the following command to start a docker container :
        ```
        ./develop_run.sh
        ```
## Running the tests

All tests run in local develoment for both docker anf Linux
If any of the test fail the full development run will fail. All details will be provided in stdout.

### Break down into end to end tests
##### There are 2 types of tests:
 - Unit tests (covering everything)
 - Integration tests (covering everything except API calls, API calls are all mocked)

## Built With

 - GitHub actions : https://github.com/AAFC-BICoE/object-store-harvestor/actions

## Artifacts
 - executable binary file
 - config yml file

## Versioning
 - Naming convention : vxxx.yyy (example : v0.4)

## Deployment
* TODO
