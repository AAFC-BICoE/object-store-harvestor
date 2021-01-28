# Design object-store-harvestor

Please refer to the [diagram](design-diagram-v0.01.pdf) in the same folder

There are 2 major blocks:
  - Local computer running on Windows or Linux where the application will walk through all media files
  - Public Internet where the Object Store API accept http requests from object-store-harvester

# Invoking object-store-harvester

  - The Application can be invoked with 2 options
    - one time run (Phase 1)
    - cron job style runs (Future phase)

# Packages of object-store-harvester
  - Orchestrater (Invokes File walker and http client) (Phase 1)
  - File walker (recursively traverse directories and find media files, connect to DB via DB client and store file paths in Local DB and mark them as 'new') (Phase 1)
  - http client (Connects to DB via DB client, get file paths with mark 'new' and try to upload a file to API and mark the file ('success' or 'fail')) (Phase 1)
  - DB client (provides connectivity to DB) (Phase 1)
  - Local DB (persistence storage) (Phase 1)
  - KeyCloak (OAuth 2 adapter for http client) (Phase 1)
  - Common Logger (provides a way to log from each component above to a file) (Phase 1)
  - Logs (it's optional for phase 1. Logs can be compressed and send to us manually for now)
