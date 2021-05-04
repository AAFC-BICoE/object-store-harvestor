# Configuration (harvestor_config.yml)

Full path to the config application file has to be passed to an application as a parametor.
Example: `./main harvestor_config.yml`
Config yml file is organized by golang packages of the application
Application will load yml file and use it for each application component respectively

## database
### maxOpenConns
`maxOpenConns:` Number of max Open connections to SQLite DB from [gorm](gorm.io) client.
### maxIdleConns
`maxIdleConns:` Number of max Idle connections to SQLite DB from [gorm](gorm.io) client.
### connMaxLifetime
`connMaxLifetime:` Number of minutes connection will live from [gorm](gorm.io) client.
### dbFile
`dbFile:` physical file path to SQLite DB. Example : `/tmp/db-test/harvestor.db`

## Walker
### maxOpenConns
`EntryPointPath:` Physical folder path for the walker recursively traverse directories and find files of application interest
### maxIdleConns
`FilesOfInterest:` comma separated list of file extensions of application interest. Example : `"yml, yaml"` or single `"yml"`

## Httpclient
### timeOut
`timeOut:` Number of seconds after client will time out after making a request to API server.
### maxIdleConnections
`maxIdleConnections:` number of http clients in the pool. For performance purposes the pool will be created at startup of the application
### retryMax
`retryMax:` number of retries with exponential backoff the http client will do before giving up
### retryWaitMin
`retryWaitMin:` minimum exponential backoff initial wait time in seconds
### baseApiUrl
`baseApiUrl:` object store base [API](https://dina-web.github.io/object-store-specs/) Url
### upload
`upload:` base upload resource on object store [API](https://dina-web.github.io/object-store-specs/) server
### uploadGroup
`uploadGroup:` uploadGroup partial resource for upload resource on object store [API](https://dina-web.github.io/object-store-specs/) server
### meta
`meta:` full metadata resource on object store [API](https://dina-web.github.io/object-store-specs/) server
### managedmeta
`managedmeta:` full managedmeta resource on object store [API](https://dina-web.github.io/object-store-specs/) server
### derivative
`derivative:` full derivative resource on object store [API](https://dina-web.github.io/object-store-specs/) server
### derivative
`derivative:` full derivative resource on object store [API](https://dina-web.github.io/object-store-specs/) server

## Keycloak
### host
`host:` host url for Keycloak
### adminClientID
`adminClientID:` The client-id of the application.
### userName
`userName:` Keycloak user name as part of credentials
### userPassword
`userPassword:` Keycloak user password as part of credentials
### grantType
`grantType:` Keycloak endpoint grant type, example: `password`
### realmName
`realmName:` realm name is a space where you manage objects. This is part of `Other realms`
### newTokenBefore
`newTokenBefore:` number of seconds before Access token expires and thhp client will be forced to get a new one 
### debug
`debug:` a boolean false or true. Very useful when you need to see full JWT token
### enabled
`enabled:` a boolean false or true. Very useful when [API](https://dina-web.github.io/object-store-specs/) server has Keycloak off (local development)

## logger
### level
`level:` logger level
### path
`level:` path to the folder where logger will log into the file
### file
`file:` file name where logger will log into

## app
### env
`env:` environment where the application will run, example `cluster` or `PC`
### objectTimezone
`objectTimezone:` EXIF does not provide timezone information so the following timezone will be used for the objects uploaded by the harvestor
